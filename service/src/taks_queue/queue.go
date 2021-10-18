package taks_queue

import (
	"runtime"
	"sync"
	"time"
)

type finalise struct {
	stop    bool
	timeout time.Duration
	msg     chan bool
}

type Handler func(interface{})

type queue struct {
	Handler          Handler
	ConcurrencyLimit uint32

	// close            time.Duration
	// killAfter        time.Duration
	finalise chan *finalise

	push      chan interface{}
	pop       chan struct{}
	suspend   chan bool
	suspended bool
	stop      chan struct{}
	stopped   bool
	buffer    []interface{}
	count     uint32
	wg        sync.WaitGroup
}

type Queue struct {
	*queue
}

func NewQueue(handler Handler, concurrencyLimit uint32) *Queue {

	q := &Queue{
		&queue{
			Handler:          handler,
			ConcurrencyLimit: concurrencyLimit,
			push:             make(chan interface{}),
			pop:              make(chan struct{}),
			suspend:          make(chan bool),
			stop:             make(chan struct{}),
		},
	}

	go q.run()
	runtime.SetFinalizer(q, (*Queue).Stop)
	return q
}

func (q *Queue) Push(val interface{}) {
	q.push <- val
}

func (q *Queue) Stop() {
	q.stop <- struct{}{}
	runtime.SetFinalizer(q, nil)
}

func (q *Queue) Wait() {
	q.wg.Wait()
}

func (q *Queue) FinaliseAndClose(timeout time.Duration) {
	result := make(chan bool)
	q.finalise <- &finalise{
		stop:    true,
		timeout: timeout,
		msg:     result,
	}

	<-result
	close(result)
}

// func (q *Queue) KillAfter(timeout time.Duration) {
// 	result := make(chan bool)
// 	throttle.msgs <- &throttleMessage{
// 		stop:   true,
// 		result: result,
// 		msg:    result,
// 	}
//
// 	<-result
// 	close(result)
// }

func (q *Queue) Len() (_, _ uint32) {
	return q.count, uint32(len(q.buffer))
}

func (q *queue) run() {

	defer func() {
		q.wg.Add(-len(q.buffer))
		q.buffer = nil
	}()

	for {

		select {

		case val := <-q.push:
			q.buffer = append(q.buffer, val)
			q.wg.Add(1)

		case <-q.pop:
			q.count--

		case suspend := <-q.suspend:
			if suspend != q.suspended {

				if suspend {
					q.wg.Add(1)
				} else {
					q.wg.Done()
				}
				q.suspended = suspend
			}

		case <-q.stop:
			q.stopped = true

		case finalise := <-q.finalise:
			time.AfterFunc(finalise.timeout, func() {
				q.buffer = []interface{}{}
				q.count = 0
				q.stopped = finalise.stop
			})
		}

		if q.stopped && q.count == 0 {
			return
		}

		for (q.count < q.ConcurrencyLimit || q.ConcurrencyLimit == 0) && len(q.buffer) > 0 && !(q.suspended || q.stopped) {
			val := q.buffer[0]
			q.buffer = q.buffer[1:]
			q.count++

			go func() {
				defer func() {
					q.pop <- struct{}{}
					q.wg.Done()
				}()

				q.Handler(val)
			}()
		}

	}

}

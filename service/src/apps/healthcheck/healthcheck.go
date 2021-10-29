package healthcheck

import (
	"log"
	"net/http"
	"os"
)

const FilePath = "/opt/healthcheck/healthy.txt"

var isDev, _ = os.LookupEnv("IS_DEV")

func isHealthy() (bool, error) {

	_, err := os.Stat(FilePath)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err

} // isHealthy

func CreateHealthCheckFile() error {

	var (
		f      *os.File
		err    error
		exists bool
	)

	if isDev == "1" {
		return err
	}

	if exists, err = isHealthy(); err != nil {
		panic(err)
	} else if exists {
		return nil
	}

	f, err = os.OpenFile(FilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte("1"))

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

	return err

} // CreateHealthCheckFile

func RemoveHealthCheckFile() {

	if isDev == "1" {
		return
	}

	if err := os.Remove(FilePath); err != nil {
		panic(err)
	}

} // RemoveHealthCheckFile

func Handler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("HEALTHY")); err != nil {
			panic(err)
		}

	}

} // Handler

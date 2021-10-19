package healthcheck

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type serviceStatus struct {
	Status string `json:"status"`
}

const FilePath = "/opt/healthcheck.txt"

var isDev, _ = os.LookupEnv("IS_DEV")

func CreateHealthCheckFile() error {
	var (
		f   *os.File
		err error
	)

	if isDev == "1" {
		return err
	}

	f, err = os.OpenFile(FilePath, os.O_CREATE|os.O_RDWR, 0650)

	if err = f.Close(); err != nil {
		panic(err)
	}

	return err
}

func RemoveHealthCheckFile() {
	if isDev == "1" {
		return
	}

	if err := os.Remove(FilePath); err != nil {
		panic(err)
	}
}

func isHealthy() (bool, error) {

	_, err := os.Stat(FilePath)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err

} // isHealthy

func Handler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			status      bool
			err         error
			jsonPayload []byte
		)

		response := &serviceStatus{}

		status, err = isHealthy()

		if err != nil || status == false {
			w.WriteHeader(http.StatusServiceUnavailable)
			response.Status = "UNHEALTHY"
		} else {
			w.WriteHeader(http.StatusOK)
			response.Status = "HEALTHY"
		}

		if jsonPayload, err = json.Marshal(response); err != nil {
			panic(err)
		} else if _, err = w.Write(jsonPayload); err != nil {
			panic(err)
		}

	}

} // Handler

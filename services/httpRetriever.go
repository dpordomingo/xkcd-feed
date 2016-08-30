package services

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func HTTPget(url string) (string, error) {
	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil || response.StatusCode != http.StatusOK {
		return "", errors.New("No resource found")
	}

	retrievedPost, err := ioutil.ReadAll(response.Body)
	return string(retrievedPost), err
}

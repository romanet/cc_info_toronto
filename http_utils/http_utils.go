package http_utils

import (
	"io"
	"log"
	"net/http"
)

func ReadURL(url string) ([]byte, error) {
	log.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, err
}

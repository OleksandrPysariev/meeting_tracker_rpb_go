package utils

import (
	"fmt"
	"io"
	"time"
	"net/http"
	"os"
)

const (
	requestURL = "http://202.61.206.185/event"
	filename = "event.json"
	maxRetries = 50
	retryDelay = 30 * time.Second
)

func ReadResponseBody() []byte {
	var response *http.Response
	var err error
	
	for attempts := 0; attempts < maxRetries; attempts++ {
		response, err = http.Get(requestURL)
		if err == nil {
			break
		}
		fmt.Printf("[utils] error making http request (attempt %d/%d): %s\n", attempts+1, maxRetries, err)
		time.Sleep(retryDelay)
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("[utils] error parsing http request body: %s\n", err)
		os.Exit(1)
    }
	return bodyBytes
}

func WriteBodyToFile(bodyBytes []byte) {
	fo, err := os.Create(filename)
	if err != nil {
		fmt.Printf("[utils] error creating file %s: %s\n", filename, err)
		os.Exit(1)
	}
	defer func() {
        if err := fo.Close(); err != nil {
            fmt.Printf("[utils] error closing file %s: %s\n", filename, err)
			os.Exit(1)
        }
    }()
	if _, err := fo.Write(bodyBytes); err != nil {
		fmt.Printf("[utils] error writing to file %s: %s\n", filename, err)
		os.Exit(1)
	}
}
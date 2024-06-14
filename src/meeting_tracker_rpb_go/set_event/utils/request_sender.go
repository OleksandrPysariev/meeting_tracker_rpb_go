package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const requestURL = "http://202.61.206.185/event"
const filename = "event.json"

func ReadResponseBody() []byte {
	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("[utils] error making http request: %s\n", err)
		os.Exit(1)
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
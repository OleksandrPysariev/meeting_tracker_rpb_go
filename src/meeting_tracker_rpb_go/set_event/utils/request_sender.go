package utils

import (
	"fmt"
	"os"
)

const (
	filename = "event.json"
)

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

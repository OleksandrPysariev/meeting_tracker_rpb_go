package read_event

import (
	"log"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PadString(str string, length int, padding string) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(padding, length-len(str))
}
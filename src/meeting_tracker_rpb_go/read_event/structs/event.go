package readjson

import "time"

type Event struct {
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Time        string `json:"time"`
	Description string `json:"description"`
}
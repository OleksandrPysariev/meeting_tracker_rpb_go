package read_event

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/OleksandrPysariev/meeting_tracker_rpb_go/read_event/structs"
)

func ParseEvent() readjson.Event {
	jsonFile, err := os.Open("event.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var event readjson.Event
	json.Unmarshal(byteValue, &event)
	return event
}
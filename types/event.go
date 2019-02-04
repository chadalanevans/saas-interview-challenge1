package types

import "time"
import "encoding/json"

type Event struct {
	Type string
	Data string
	Created time.Time
}

func (event Event) String() string {
	var eventStr, _ = json.Marshal(event)
	return string(eventStr)
}

func (event *Event) FromString(data string) error {
	return json.Unmarshal([]byte(data), event)	
}
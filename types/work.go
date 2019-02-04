package types

import "encoding/json"

type Work struct {
	Id uint64
	Type string
	State WorkState
	Data string
}

func (work Work) String() string {
	var workStr, _ = json.Marshal(work)
	return string(workStr)
}

func (work *Work) FromString(data string) error {
	return json.Unmarshal([]byte(data), work)	
}
package messagetype

import "encoding/json"

type Type struct {
	Type int `json:"type"`
}

func ParseInput(b []byte) InputType {
	var obj Type
	_ = json.Unmarshal(b, &obj)
	return InputType(obj.Type)
}

type InputType int

const SendMessage InputType = iota + 1

var SentMessage = Type{1}

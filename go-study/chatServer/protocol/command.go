package protocol

import "errors"

var (
	UnKnowCommand = errors.New("unKnow  command...")
)
type SendCommand struct{
	Message string
}

type NameCommand struct {
	Name string
}

type MessCommand struct {
	Name string
	Message string
}
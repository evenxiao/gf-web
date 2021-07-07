package protocol

import (
	"fmt"
	"io"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {

	return &Writer{
		writer:w,
	}
}

func (w *Writer) WriteString(msg string) error{
	_, err := w.writer.Write([]byte(msg))
	return err
}

func (w *Writer) Write(command interface{}) error {
	var err error

	switch v := command.(type) {
	case SendCommand:
		err = w.WriteString(fmt.Sprintf("Command : SEND %v \n", v.Message))
	case NameCommand:
		err = w.WriteString(fmt.Sprintf("Command : NAME %v \n", v.Name))

	case MessCommand:
		err = w.WriteString(fmt.Sprintf("Command:Message %v %v", v.Name, v.Message))

	default:
		err = UnKnowCommand
	}

	return err
}

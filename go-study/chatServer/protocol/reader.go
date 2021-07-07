package protocol

import (
	"bufio"
	"io"
)

type Reader struct {
	reader *bufio.Reader
}

func NewReader(reader io.Reader) *Reader {

	return &Reader{
		reader:bufio.NewReader(reader),
	}
}

func (r *Reader) Read() (interface{}, error)   {

	cmd, err := r.reader.ReadString( ' ' )

	if err != nil {
		return nil, err
	}
	switch cmd {

	case "MESS":
		user, err := r.reader.ReadString(' ')

		if err != nil {
			 return nil, err
		}

		message, err := r.reader.ReadString('\n')

		if err  != nil{
			return nil, err
		}
		return MessCommand{
			user[:len(user)-1],
			message[:len(message)-1],
		}, nil
	}

	return nil, UnKnowCommand
}
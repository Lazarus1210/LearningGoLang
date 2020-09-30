package protocol

import (
	"fmt"
	"io"
)

// CommandWriter is a struct that has a variable of type io.Writer
type CommandWriter struct {
	writer io.Writer
}

// NewCommandWriter take io.Writer as an argument
func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{
		writer: writer,
	}
}

func (w *CommandWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))

	return err
}

func (w *CommandWriter) Write(command interface{}) error {
	var err error
	switch v := command.(type) {
	case SendCommand:
		err = w.writeString(fmt.Sprintf("SEND %v\n", v.command))
	case NameCommand:
		err = w.writeString(fmt.Sprintf("NAME %v\n", v.clientName))
	case MessageCommand:
		err = w.writeString(fmt.Sprintf("MESSAGE %v %v\n", v.clientName, v.message))
	default:
		err = errUnknownCommand
	}
	return err
}

package protocol

import "errors"

// Unknown Command
var (
	errUnknownCommand = errors.New("Unknown Command from the client")
)

// SendCommand is used by clients to send a new command
type SendCommand struct {
	command string
}

// NameCommand is used by client to send its name
type NameCommand struct {
	clientName string
}

// MessageCommand is the actual message send by the client
type MessageCommand struct {
	clientName string
	message    string
}

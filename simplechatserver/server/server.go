package server

// ChatServer interface, what all things should a server do
type ChatServer interface {
	Listen(address string) error
	BroadCast(command interface{}) error
	Start()
	Close()
}

package specialuse

import (
	"github.com/sebas05000/go-imap/server"
)

type extension struct{}

func (ext *extension) Capabilities(c server.Conn) []string {
	return nil
}

func (ext *extension) Command(name string) server.HandlerFactory {
	return nil
}

func NewExtension() server.Extension {
	return &extension{}
}

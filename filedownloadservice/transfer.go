package filedownloadservice

import (
	"context"
	"filedownloadservice/mybps"
	"io"
)

type transfer struct {
	n     int64 // must be 64bit aligned on 386
	ctx   context.Context
	gauge mybps.Guage
	lim   RateLimiter
	w     io.Writer
	r     io.Reader
	b     []byte
}

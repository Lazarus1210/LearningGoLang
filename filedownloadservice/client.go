package filedownloadservice

import (
	"context"
	"net/http"
	"time"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// the client struct is a file download client, this has some other properties as well on top of http.Client
type myClient struct {
	//variable of the interface HTTPClient type
	httpClient HTTPClient
	// the user agent string to be passed in requests
	userAgent string
	//the buffer size to transfer the requested files
	bufferSize int
}

// this will return a new client instance
func newClient() *myClient {
	return &myClient{
		userAgent: "filedownloadservice",
		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
	}
}

// default client that will be used will default properties
var defaultClient = newClient()

// Do sends a file transfer request and returns a file transfer response,
// following policy (e.g. redirects, cookies, auth) as configured on the
// client's HTTPClient.
//
// Like http.Get, Do blocks while the transfer is initiated, but returns as soon
// as the transfer has started transferring in a background goroutine, or if it
// failed early.
//
// An error is returned via Response.Err if caused by client policy (such as
// CheckRedirect), or if there was an HTTP protocol or IO error. Response.Err
// will block the caller until the transfer is completed, successfully or
// otherwise.

func (c *myClient) Do(req *myRequest) myResponse {
	ctx, cancel := context.WithCancel(req.Context())
	resp := &myResponse{
		Request:    req,
		Start:      time.Now(),
		Filename:   req.Filename,
		ctx:        ctx,
		cancel:     cancel,
		bufferSize: req.BufferSize,
	}

	if resp.bufferSize == 0 {
		resp.bufferSize = c.bufferSize
	}

}

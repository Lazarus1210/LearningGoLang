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

// An stateFunc is an action that mutates the state of a Response and returns
// the next stateFunc to be called.
// Learning : this is one of the usage of closure. This function take myResponse as the argument and returns another
// state function, thus forming a series of actions that can be taken.
// Also, all the functions implement myClient struct, hence all these acceprt myClient as a reciever and are member functions of myclient struct. These can be inviked by any variable of of type myclient struct

type stateFunc func(*myResponse) stateFunc

// run calls the given stateFunc function and all subsequent returned stateFuncs
// until a stateFunc returns nil or the Response.ctx is canceled.
// Implements myClient, i.e takes this as a reciever. Can be invoked by any variable of myClient type
// Each stateFunc
// should mutate the state of the given Response until it has completed
// downloading or failed.

func (c *myClient) run(resp *myResponse, f stateFunc) stateFunc {
	for {
		select {
		case <-resp.ctx.Done():
			if resp.IsComplete() {
				return
			}
			resp.err = resp.ctx.Err()

		}
	}
}

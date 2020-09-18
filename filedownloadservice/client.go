package filedownloadservice

import "net/http"

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

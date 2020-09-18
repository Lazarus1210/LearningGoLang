package filedownloadservice

import (
	"context"
	"hash"
	"net/http"
)

// A hook is a user provided callback to track the progress in the response. These can be called by filedownload service at any point in when the request is recieved
type userHook func(*myResponse) error

type myRequest struct {

	// Label is an arbitrary string which may used to label a Request with a
	// user friendly name.
	Label string

	// Tag is an arbitrary interface which may be used to relate a Request to
	// other data.
	Tag interface{}

	// HTTPRequest specifies the http.Request to be sent to the remote server to
	// initiate a file transfer. It includes request configuration such as URL,
	// protocol version, HTTP method, request headers and authentication.
	HTTPRequest *http.Request

	// Filename specifies the path where the file transfer will be stored in
	// local storage. If Filename is empty or a directory, the true Filename will
	// be resolved using Content-Disposition headers or the request URL.
	//
	// An empty string means the transfer will be stored in the current working
	// directory.
	Filename string

	// SkipExisting specifies that ErrFileExists should be returned if the
	// destination path already exists. The existing file will not be checked for
	// completeness.
	SkipExisting bool

	// NoResume specifies that a partially completed download will be restarted
	// without attempting to resume any existing file. If the download is already
	// completed in full, it will not be restarted.
	NoResume bool

	// NoStore specifies that grab should not write to the local file system.
	// Instead, the download will be stored in memory and accessible only via
	// Response.Open or Response.Bytes.
	NoStore bool

	// NoCreateDirectories specifies that any missing directories in the given
	// Filename path should not be created automatically, if they do not already
	// exist.
	NoCreateDirectories bool

	// IgnoreBadStatusCodes specifies that grab should accept any status code in
	// the response from the remote server. Otherwise, grab expects the response
	// status code to be within the 2XX range (after following redirects).
	IgnoreBadStatusCodes bool

	// IgnoreRemoteTime specifies that grab should not attempt to set the
	// timestamp of the local file to match the remote file.
	IgnoreRemoteTime bool

	// Size specifies the expected size of the file transfer if known. If the
	// server response size does not match, the transfer is cancelled and
	// ErrBadLength returned.
	Size int64

	// BufferSize specifies the size in bytes of the buffer that is used for
	// transferring the requested file. Larger buffers may result in faster
	// throughput but will use more memory and result in less frequent updates
	// to the transfer progress statistics. If a RateLimiter is configured,
	// BufferSize should be much lower than the rate limit. Default: 32KB.
	BufferSize int

	// RateLimiter allows the transfer rate of a download to be limited. The given
	// Request.BufferSize determines how frequently the RateLimiter will be
	// polled.
	RateLimiter RateLimiter

	// BeforeCopy is a user provided callback that is called immediately before
	// a request starts downloading. If BeforeCopy returns an error, the request
	// is cancelled and the same error is returned on the Response object.
	BeforeCopy userHook

	// AfterCopy is a user provided callback that is called immediately after a
	// request has finished downloading, before checksum validation and closure.
	// This hook is only called if the transfer was successful. If AfterCopy
	// returns an error, the request is canceled and the same error is returned on
	// the Response object.
	AfterCopy userHook

	// hash, checksum and deleteOnError - set via SetChecksum.
	hash          hash.Hash
	checksum      []byte
	deleteOnError bool

	// Context for cancellation and timeout - set via WithContext
	ctx context.Context
}

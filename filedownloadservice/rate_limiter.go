package filedownloadservice

import "context"

// RateLimiter is an interface that must be satisfied by any third-party rate
// limiters that may be used to limit download transfer speeds.
type RateLimiter interface {
	WaitN(ctx context.Context, n int) (err error)
}

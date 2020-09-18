/*
Package bps provides gauges for calculating the Bytes Per Second transfer rate
of data streams.
*/

package mybps

import "time"

// Gauge is the common interface for all BPS gauges in this package. Given a
// set of samples over time, each gauge type can be used to measure the Bytes
// Per Second transfer rate of a data stream.
//
// All samples must monotonically increase in timestamp and value. Each sample
// should represent the total number of bytes sent in a stream, rather than
// accounting for the number sent since the last sample.
//
// To ensure a gauge can report progress as quickly as possible, take an initial
// sample when your stream first starts.
//
// All gauge implementations are safe for concurrent use.
type Gauge interface {
	// Sample adds a new sample of the progress of the monitored stream.
	Sample(t time.Time, n int64)

	// BPS returns the calculated Bytes Per Second rate of the monitored stream.
	BPS() float64
}

// This file contains set of Go functions that handle backoff strategies for HTTP clients. These strategies are typically used when a client makes a request to a server and, in case of failure, determines how long to wait before trying again.
package client

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"net/http"
	"sync"
	"time"
)

// Backoff specifies a policy for how long to wait between retries.
type Backoff func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration

// DefaultBackoff provides a callback for client.Backoff
// implements the standard exponential backoff without jitter.
// i.e The delay between retries is doubled with each attempt, up to a maximum delay.
func DefaultBackoff() func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	return func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		mult := math.Pow(2, float64(attemptNum)) * float64(min)

		sleep := time.Duration(mult)

		if float64(sleep) != mult || sleep > max {
			sleep = max
		}

		return sleep
	}
}

// LinearJitterBackoff provides a callback for client.Backoff which
// implements linear backoff with jitter.
// i.e The delay between retries is increased linearly with each attempt,
// but a random jitter is added to this delay.
//
// This jitter helps in distributed systems to avoid situations
// where many clients retry simultaneously, commonly known as "thundering herd".
//
// min and max here are *not* absolute values. The number to be multiplied by
// the attempt number will be chosen at random from between them, thus they are
// bounding the jitter.
//
// For instance:
// - To get strictly linear backoff of one second increasing each retry, set
// both to one second (1s, 2s, 3s, 4s, ...)
// - To get a small amount of jitter centered around one second increasing each
// retry, set to around one second, such as a min of 800ms and max of 1200ms
// (892ms, 2102ms, 2945ms, 4312ms, ...)
// - To get extreme jitter, set to a very wide spread, such as a min of 100ms
// and a max of 20s (15382ms, 292ms, 51321ms, 35234ms, ...)
func LinearJitterBackoff() func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	randMutex := &sync.Mutex{}

	return func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		// attemptNum always starts at zero but we want to start at 1 for multiplication
		attemptNum++

		if max <= min {
			// Unclear what to do here, or they are the same, so return min *
			// attemptNum
			return min * time.Duration(attemptNum)
		}

		// Pick a random number that lies somewhere between the min and max and
		// multiply by the attemptNum. attemptNum starts at zero so we always
		// increment here. We first get a random percentage, then apply that to the
		// difference between min and max, and add to min.
		randMutex.Lock()
		jitter := cryptoRandFloat64() * float64(max-min)
		randMutex.Unlock()

		jitterMin := int64(jitter) + int64(min)

		return time.Duration(jitterMin * int64(attemptNum))
	}
}

// FullJitterBackoff provides a callback for client.Backoff which
// implements a variation of exponential backoff with full jitter.
// i.e Instead of doubling the delay with each attempt, it randomizes the delay completely within the exponential window.
//
// Algorithm is fast because it does not use floating
// point arithmetics. It returns a random number between [0...n]
// https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
func FullJitterBackoff() func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	randMutex := &sync.Mutex{}

	return func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		duration := attemptNum * 1000000000 << 1

		randMutex.Lock()
		jitter := cryptoRandInt(duration-attemptNum) + int(min)
		randMutex.Unlock()

		if jitter > int(max) {
			return max
		}

		return time.Duration(jitter)
	}
}

// ExponentialJitterBackoff provides a callback for Client.Backoff which will
// perform an exponential backoff based on the attempt number and with jitter to
// prevent a thundering herd.
//
// min and max here are *not* absolute values. The number to be multiplied by
// the attempt number will be chosen at random from between them, thus they are
// bounding the jitter.
func ExponentialJitterBackoff() func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	randMutex := &sync.Mutex{}

	return func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		minf := float64(min)
		mult := math.Pow(2, float64(attemptNum)) * minf

		randMutex.Lock()
		jitter := cryptoRandFloat64() * (mult - minf)
		randMutex.Unlock()

		mult += jitter

		sleep := time.Duration(mult)

		if sleep > max {
			sleep = max
		}

		return sleep
	}
}

// Helper function to get a float64 value between 0 and 1 using crypto/rand
func cryptoRandFloat64() float64 {
	var buf [8]byte

	_, err := rand.Read(buf[:])
	if err != nil {
		panic(err) // handle this error appropriately
	}

	// Convert the uint64 to a float64 in [0, 1)
	return float64(binary.LittleEndian.Uint64(buf[:])) / float64(1<<64)
}

// Helper function to get a random integer between 0 and max using crypto/rand
func cryptoRandInt(max int) int {
	if max <= 0 {
		return 0
	}

	var n uint64

	max64 := uint64(max)
	buf := make([]byte, 8)

	for {
		_, err := rand.Read(buf)
		if err != nil {
			panic(err) // handle this error appropriately
		}

		n = binary.LittleEndian.Uint64(buf)

		if n < (1<<63 - (1 << 63 % max64)) {
			return int(n % max64)
		}
	}
}

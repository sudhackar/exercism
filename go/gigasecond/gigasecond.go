// Package gigasecond provides a function to add a gigasec(10**9)
// to a given time
package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

// AddGigasecond returns a time with added gigasecond to t
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(math.Pow10(9)))
}

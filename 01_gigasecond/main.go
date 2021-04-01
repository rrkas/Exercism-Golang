// Given a moment, determine the moment that would be after a gigasecond has passed.
package main

// import path for the time package from the standard library
import (
	"fmt"
	"time"
)

// Add one gigasecond (10^9 seconds) to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}

func main() {
	t := time.Now()
	fmt.Print(AddGigasecond(t))
}

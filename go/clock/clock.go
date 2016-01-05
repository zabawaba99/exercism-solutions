package clock

import "fmt"

// The value of TestVersion here must match `testVersion` in the file
// clock_test.go.
const TestVersion = 2

const (
	maxHour = 24
	maxTime = Clock(60 * 24)
	minTime = Clock(0)
)

// Clock represents a time value only. It does not
// take dates into account.
type Clock int

// Time creates a new clock object
func Time(hour, minute int) (c Clock) {
	// ensure that we do not go over the max time
	hour %= maxHour

	t := hour*60 + minute
	return c.Add(t)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add returns the clock c+minutes
func (c Clock) Add(minutes int) Clock {
	c2 := c + Clock(minutes)

	// if we have a negative number, we need to
	// roll back
	if c2 < minTime {
		c2 += maxTime
	}

	// if our number is greater than the max number
	// of minutes, we need to roll over
	if c2 > maxTime {
		c2 -= maxTime
	}
	return c2
}

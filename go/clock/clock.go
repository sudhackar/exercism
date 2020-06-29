package clock

import "fmt"

// Clock wraps hours and minutes
type Clock struct {
	hours, mins int
}

func normalize(c Clock) Clock {
	if c.mins > 59 || c.mins < 0 {
		c.hours += (c.mins / 60)
		c.mins %= 60
		if c.mins < 0 {
			c.mins += 60
			c.hours--
		}
	}
	c.hours %= 24
	c.hours += 24
	c.hours %= 24

	return c
}

// New creates a new Clock
func New(hours, mins int) Clock {
	c := new(Clock)
	c.hours = hours
	c.mins = mins
	return normalize(*c)
}

// Add mins to Clock
func (c Clock) Add(mins int) Clock {
	return normalize(Clock{
		c.hours,
		c.mins + mins,
	})
}

// Subtract mins from Clock
func (c Clock) Subtract(mins int) Clock {
	return normalize(Clock{
		c.hours,
		c.mins - mins,
	})
}

// Strring returns a string representation of Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.mins)
}

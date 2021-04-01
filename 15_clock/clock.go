package clock

import "fmt"

const minutesInADay = 24 * 60

type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	minutes := ((60*h+m)%minutesInADay + minutesInADay) % minutesInADay

	return Clock{minutes / 60, minutes % 60}
}

func (c Clock) Add(min int) Clock {
	minutes := ((60*c.h+c.m+min)%minutesInADay + minutesInADay) % minutesInADay

	return Clock{minutes / 60, minutes % 60}
}

func (c Clock) Subtract(min int) Clock {
	minutes := ((60*c.h+c.m-min)%minutesInADay + minutesInADay) % minutesInADay

	return Clock{minutes / 60, minutes % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

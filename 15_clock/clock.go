package clock

import "fmt"

type Clock struct {
	h, m int
}

func (c Clock) Add(m int) Clock {
	c.m += m
	c.h += c.m / 60
	c.m %= 60
	if c.h >= 24 || c.m >= 60 {
		c.h += c.m / 60
		c.h %= 24
		c.m %= 60
	}
	return Clock{c.h, c.m}
}

func (c Clock) Subtract(m int) Clock {
	c.m -= m
	c.h += c.m / 60
	c.m %= 60
	if c.m < 0 {
		c.m %= 60
		c.m = 60 + c.m
		c.h--
	}
	if c.h < 0 {
		c.h %= 24
		c.h = 24 + c.h
	}
	if c.h >= 24 || c.m >= 60 {
		c.h += c.m / 60
		c.h %= 24
		c.m %= 60
	}
	return Clock{c.h, c.m}
}

func (c Clock) String() string {
	s := ""
	if c.h < 10 {
		s = s + "0"
	}
	s = s + fmt.Sprint(c.h)
	s = s + ":"
	if c.m < 10 {
		s = s + "0"
	}
	s = s + fmt.Sprint(c.m)
	return s
}

func New(h, m int) Clock {
	if h < 0 || m < 0 {
		if m < 0 {
			h += m / 60
			m %= 60
		}
		if h < 0 {
			h %= 24
			h = 24 + h
		}
		if m < 0 {
			h--
			m = 60 + m
		}
	}
	if h >= 24 || m >= 60 {
		h += m / 60
		h %= 24
		m %= 60
	}
	return Clock{h, m}
}

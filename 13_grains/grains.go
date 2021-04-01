package grains

import (
	"errors"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("error")
	}
	var s uint64 = 1
	for i := 1; i < n; i++ {
		s *= 2
	}
	return s, nil
}

func Total() uint64 {
	var s uint64 = 0
	for i := 1; i < 65; i++ {
		t, _ := Square(int(i))
		s += t
	}
	return s
}

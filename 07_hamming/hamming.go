package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("error")
	}
	c := 0
	for i := 0; i < len(a); i = i + 1 {
		if a[i] != b[i] {
			c = c + 1
		}
	}
	return c, nil
}

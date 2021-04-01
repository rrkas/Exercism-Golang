package reverse

func Reverse(text string) string {
	s := ""
	for _, c := range text {
		s = string(c) + s
	}
	return s
}

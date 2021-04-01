package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(sl []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap, len(sl))
	for _, s := range sl {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for i := 0; i < len(sl); i++ {
		for k, v := range <-c {
			m[k] += v
		}
	}

	return m
}

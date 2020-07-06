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

// ConcurrentFrequency counts the frequency of each rune in a given text concurrently
// and returns this data as a FreqMap.
func ConcurrentFrequency(words []string) FreqMap {
	res := FreqMap{}
	channel := make(chan FreqMap, len(words))

	for _, word := range words {
		go func(str string) {
			channel <- Frequency(str)
		}(word)
	}

	for range words {
		for c, n := range <-channel {
			res[c] += n
		}
	}
	return res
}

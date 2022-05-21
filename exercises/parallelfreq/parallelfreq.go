package parallelfreq

type FreqMap map[rune]int

// Count frequency of characters in a string using concurrency

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(l []string) FreqMap {
    c := make(chan rune)
    m := make(FreqMap)
	count := 0
	signal := 0

    for _, str := range l {
		count += len(str)
        go sendUpdateCount(str, c)
    }
	
    for char := range c {
        m[char]++
		signal++
		if signal == count {
			close(c)
		}
    }
	
    return m
}

func sendUpdateCount(s string, c chan rune) {
	n := len(s)
	i := 0
    for i < n {
		char := rune(s[i])
		select {
			case c <- char:
				i++
		}
        
	}

}

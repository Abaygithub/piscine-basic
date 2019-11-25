package piscine

func AlphaCount(str string) int {
	index := 0
	for _, letter := range str {
		if letter >= 65 && letter <= 90 || letter >= 97 && letter <= 122 {
			index++
		}

	}
	return index
}

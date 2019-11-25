package piscine

func Index(s string, toFind string) int {
	myArray := []rune(s)
	incl := []rune(toFind)
	len1 := 0
	len2 := 0

	for range myArray {
		len1++
	}

	for range incl {
		len2++
	}

	for i := 0; i <= len1-len2; i++ {
		if toFind == s[i:i+len2] {
			return i
		}
	}
	return -1
}

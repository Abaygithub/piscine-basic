package piscine

func ToLower(s string) string {
	var result string
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			letter = letter + 32
			result = result + string(letter)
		} else {
			result = result + string(letter)
		}
	}
	return result
}

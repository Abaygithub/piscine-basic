package piscine

func ToUpper(s string) string {
	var result string
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			letter = letter - 32
			result = result + string(letter)
		} else {
			result = result + string(letter)
		}
	}
	return result

}

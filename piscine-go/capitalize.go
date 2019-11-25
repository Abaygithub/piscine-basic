package piscine

func Capitalize(s string) string {

	newWord := true
	myRune := []rune(s)
	for index, value := range myRune {
		if ('a' <= value && value <= 'z') && newWord {
			value = value - 32
			myRune[index] = value
			newWord = false
		} else if ('A' <= value && value <= 'Z') && !newWord {
			value = value + 32
			myRune[index] = value
		}
		if !('a' <= value && value <= 'z') && !('A' <= value && value <= 'Z') && !('1' <= value && value <= '9') {
			newWord = true
		} else {
			newWord = false
		}
	}
	return string(myRune)
}

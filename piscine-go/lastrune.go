package piscine

func LastRune(s string) rune {
	myArray := []rune(s)
	var ru rune
	for _, l := range myArray {
		ru = l
	}
	return ru
}

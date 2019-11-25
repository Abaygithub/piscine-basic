package piscine

func FirstRune(s string) rune {
	array := []rune(s)
	var r rune
	for i, v := range array {
		if i == 0 {
			r = v
		}
	}
	return r
}

package piscine

func NRune(s string, n int) rune {
	array := []rune(s)
	var t rune
	if n > 0 {
		for i, l := range array {
			if i == n-1 {
				t = l
			}
		}
	}
	return t
}

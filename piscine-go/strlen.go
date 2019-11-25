package piscine

func StrLen(str string) int {
	r := []rune(str)
	ln := 0
	for l := range r {
		ln = l
	}
	ln = ln + 1
	return ln
}

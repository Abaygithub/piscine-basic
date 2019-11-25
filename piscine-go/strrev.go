package piscine

func StrRev(s string) string {
	reverse := []rune(s)
	var a string
	for _, ran := range reverse {
		a = string(ran) + a
	}
	return a
}

package piscine

func IsUpper(str string) bool {
	n := true
	for _, i := range str {
		if !(i >= 'A' && i <= 'Z') {
			n = false
		}
	}
	return n

}

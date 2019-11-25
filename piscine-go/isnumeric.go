package piscine

func IsNumeric(str string) bool {
	num := true
	for _, i := range str {
		if !(i >= '0' && i <= '9') {
			num = false
		}
	}
	return num
}

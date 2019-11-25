package piscine

func IsPrintable(str string) bool {
	flag := true
	for _, i := range str {
		if i >= 0 && i <= 31 {
			flag = false
		}
	}
	return flag
}

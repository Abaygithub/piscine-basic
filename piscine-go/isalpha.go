package piscine

func IsAlpha(str string) bool {
	alpha := true
	for _, i := range str {
		if !(i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' || i >= '0' && i <= '9') {
			alpha = false
		}
	}
	return alpha

}

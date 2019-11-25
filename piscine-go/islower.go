package piscine

func IsLower(str string) bool {
	nu := true
	for _, i := range str {
		if !(i >= 'a' && i <= 'z') {
			nu = false
		}
	}
	return nu
}

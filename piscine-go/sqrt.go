package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb <= 0 {
		return 0
	}
	for s := 1; s <= nb/2; s++ {
		if s*s == nb {
			return s
		}
	}
	return 0
}

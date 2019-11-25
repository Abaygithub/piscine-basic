package piscine

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	if nb <= 0 {
		return false
	} else {
		for s := 2; s*s <= nb; s++ {
			if nb%s == 0 {
				return false
			}
		}
		return true
	}
}

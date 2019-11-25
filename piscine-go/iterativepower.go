package piscine

func IterativePower(nb int, power int) int {
	var output int = 1
	if power < 0 {
		return 0
	} else {
		for i := 1; i <= power; i++ {
			output *= nb
		}
		return output
	}
	return 0
}

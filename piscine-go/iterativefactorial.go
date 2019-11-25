package piscine

func IterativeFactorial(nb int) int {
	output := 1
	if nb >= 0 && nb < 21 {
		for i := 1; i <= nb; i++ {
			output = output * i
		}
		return output
	} else {
		return 0
	}
}

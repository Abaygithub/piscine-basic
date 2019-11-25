package piscine

func MakeRange(min, max int) []int {
	if max > min {
		answer := make([]int, max-min)
		var myVar int
		for i := min; i < max; i++ {
			answer[myVar] = i
			myVar++
		}
		return answer
	}
	return nil
}

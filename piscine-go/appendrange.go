package piscine

func AppendRange(min, max int) []int {
	var out []int
	for in := min; in < max; in++ {
		out = append(out, in)
	}
	return out
}

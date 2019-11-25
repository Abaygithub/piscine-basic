package piscine

func ConcatParams(args []string) string {
	answer := ""
	for i, v := range args {
		if i == 0 {
			answer = v
			continue
		}
		answer = answer + "\n" + v
	}
	return answer
}

package piscine

func len(input string) int {
	lengh := 0
	for range input {
		lengh++
	}
	return lengh
}

func Split(str, charset string) []string {

	var end string
	end += str + charset
	r1 := []rune(end)
	l1 := len(string(r1))
	r2 := []rune(charset)
	l2 := len(string(r2))
	nbword := 0
	for i := 0; i <= l1-l2; i++ {
		if end[i:i+l2] == charset {
			nbword++
		}
	}

	arr := make([]string, nbword)
	index := 0
	j := 0
	for i := index; i <= l1-l2; i++ {
		if end[i:i+l2] == charset {
			arr[j] = end[index:i]
			index = i + l2
			j++
		}
	}
	return arr
}

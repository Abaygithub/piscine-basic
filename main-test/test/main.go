package main

import "github.com/01-edu/z01"

func SplitWhiteSpaces(str string) []string {
	var len, index int
	for i, j := range str {
		if (j == '\n' || j == '\t' || j == ' ') && (i != 0 && str[i-1] != '\n' && str[i-1] != '\t' && str[i-1] != ' ') {
			len = len + 1
		}
	}

	var words string
	answer := make([]string, len+1)
	for _, j := range str {
		if j != ' ' && j != '\n' && j != '\t' {
			words += string(j)
		} else if (j == ' ' || j == '\n' || j == '\t') && words != "" {
			answer[index] = words
			words, index = "", index+1
		}
	}
	answer[index] = words
	return answer
}

func PrintWordsTables(table []string) {
	for i := range table {
		myRunes := []rune(table[i])
		for si := range myRunes {
			z01.PrintRune(myRunes[si])
		}
		z01.PrintRune('\n')
	}
}

func main() {
	str := "Hello how are you?"
	table := SplitWhiteSpaces(str)
	PrintWordsTables(table)
}

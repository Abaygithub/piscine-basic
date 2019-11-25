package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for i := range table {
		myRunes := []rune(table[i])
		for si := range myRunes {
			z01.PrintRune(myRunes[si])
		}
		z01.PrintRune('\n')
	}
}

package main

import "github.com/01-edu/z01"

func main() {

	num := 122

	for num >= 97 {

		z01.PrintRune(rune(num))

		num--

	}

	z01.PrintRune('\n')
}

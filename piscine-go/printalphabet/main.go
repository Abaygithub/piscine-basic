package main

import "github.com/01-edu/z01"

func main() {

	num := 97

	for num <= 122 {

		z01.PrintRune(rune(num))

		num++

	}

	z01.PrintRune('\n')
}

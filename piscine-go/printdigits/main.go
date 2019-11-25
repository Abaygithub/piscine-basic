package main

import "github.com/01-edu/z01"

func main() {

	num := '0'

	for num <= '9' {

		z01.PrintRune(num)

		num++

	}

	z01.PrintRune('\n')
}

package main

import "github.com/01-edu/z01"

func main() {

	for ascii := 33; ascii < 127; ascii++ {
		z01.PrintRune(rune(ascii))
	}
}

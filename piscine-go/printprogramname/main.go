package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myArg := []rune(os.Args[0])
	for _, v := range myArg {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

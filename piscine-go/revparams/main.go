package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	lenArg := 0
	for index := range arguments {
		lenArg = index
	}

	for i := lenArg; i >= 1; i-- {
		for _, value := range arguments[i] {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}

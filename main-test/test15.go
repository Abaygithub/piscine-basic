package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	str := "Hello"

	fmt.Println("1 " + str)

	myRune := []rune(str)
	fmt.Print("2 ")
	fmt.Print(myRune)
	z01.PrintRune('\n')

	fmt.Print("3 ")
	for x := range myRune {
		z01.PrintRune(rune(myRune[x]))
	}
	z01.PrintRune('\n')
}

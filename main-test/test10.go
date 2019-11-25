package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	var myVar string
	myVar = "Hello! How are you? How+are+things+4you?"
	myRune := []rune(myVar)

	fmt.Println(myRune)
	z01.PrintRune("myRune")
	z01.PrintRune('2')
	z01.PrintRune('\n')
}

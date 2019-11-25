package main

import "github.com/01-edu/z01"

func PrintStr(str string) {
	loop := []byte(str)
	// for x := range loop {
		z01.PrintRune(loop)
	}


func main() {
	str := "Hello World!"
	PrintStr(str)
}

package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	loop := []byte(str)
	for x := range loop {
		z01.PrintRune(rune(loop[x]))
	}
}

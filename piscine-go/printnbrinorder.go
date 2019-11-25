package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var myRune rune
	var myString string

	if n <= 0 {
		z01.PrintRune(48)
	} else {
		for n > 0 {
			j := (n % 10) + 48 //превращяем в руны
			n = n / 10
			myRune = rune(j)
			myString = myString + string(myRune) //записываем в строку
		}
		outRune := []rune(myString)
		for i := '0'; i <= '9'; i++ { //сортируем массив
			for _, digits := range outRune {
				if i == digits {
					z01.PrintRune(digits)
				}
			}
		}
	}
}

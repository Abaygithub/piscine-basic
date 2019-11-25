package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	// lenArg := 0
	// for index := range arguments {
	// 	lenArg = index
	// }

	// var myOrder string
	// for ascii := 33; ascii < 127; ascii++ {
	// 	for
	// 	for _, value := range arguments {
	// 		if ascii == value {
	// 			myOrder = myOrder + value
	// 			fmt.Println(myOrder)
	// 		}

	// 	}
	// }

	// for i := 0; i < lenArg; i++ {
	// 	for j := 0; j < lenArg-1-i; j++ {
	// 		if arguments[j] > arguments[j+1] {
	// 			arguments[j], arguments[j+1] = arguments[j+1], arguments[j]
	// 		}
	// 	}
	// }

	// for a := range arguments {
	// 	for _, x := range arguments[a] {
	// 		z01.PrintRune(x)
	// 	}
	// 	z01.PrintRune('\n')
	// }

	// for a := range arguments {
	// 	for _, x := range arguments[a] {
	// 		for ascii := 33; ascii < 127; ascii++ {
	// 			if x == rune(ascii) {
	// 				z01.PrintRune(x)
	// 			}
	// 		}
	// 	}
	// 	z01.PrintRune('\n')
	// }

	// for a := range arguments {
	// 	for ascii := 33; ascii < 127; ascii++ {
	// 		for _, x := range arguments[a] {
	// 			if rune(ascii) == x {
	// 				z01.PrintRune(x)
	// 			}
	// 		}
	// 	}
	// 	z01.PrintRune('\n')
	// }

	// for a := range arguments {
	// 	for _, x := range arguments[a] {
	// 		for ascii := 33; ascii < 127; ascii++ {
	// 			if x == rune(ascii) {
	// 				z01.PrintRune(x)
	// 			}
	// 		}
	// 	}
	// }
	// z01.PrintRune('\n')

	// for a := range arguments {
	// 	for ascii := 33; ascii < 127; ascii++ {
	// 		for _, x := range arguments[a] {
	// 			if rune(ascii) == x {
	// 				z01.PrintRune(x)
	// 			}
	// 		}
	// 	}
	// 	z01.PrintRune('\n')
	// }

	for a := range arguments {
		for ascii := 33; ascii < 127; ascii++ {
			for _, x := range arguments[a] {
				if x == rune(ascii) {
					z01.PrintRune(x)
				}
			}

		}

	}
	z01.PrintRune('\n')
}

package main

import "github.com/01-edu/z01"

func Raid1a(x,y int) {
	for line := 1; line <= x; line++ {
		for column := 1; column <= y; column++ {
			if line == 1 || line == x {
				if column == 1 || column == y {
					z01.PrintRune(111)
				} else {
					z01.PrintRune(45)
				}
			} else {
				if column == 1 || column == y {
					z01.PrintRune(124)
				} else {
					z01.PrintRune(95)
				}
			}
		}
		z01.PrintRune('\n')
	}

}


func main() {
	Raid1a(5,5)
}

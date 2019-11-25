package main

import "github.com/01-edu/z01"

func Raid1b(x, y int) {
	for line := 1; line <= y; line++ {
		for column := 1; column <= x; column++ {
			if line == 1 && column == 1 {
				z01.PrintRune('/') ///
			} else if (line == 1 && column == x) || (column == 1 && line == y) {
				z01.PrintRune('\\') //\
			} else if column == x && line == y {
				z01.PrintRune('/') ///
			} else if column == 1 || column == x || line == 1 || line == y {
				z01.PrintRune('*') ///
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	Raid1b(5, 3)
	z01.PrintRune('\n')
	Raid1b(5, 1)
	z01.PrintRune('\n')
	Raid1b(1, 1)
	z01.PrintRune('\n')
	Raid1b(1, 5)
}

package main

import "github.com/01-edu/z01"

func Raid1c(x, y int) {
	for line := 1; line <= y; line++ {
		for column := 1; column <= x; column++ {
			if (line == 1 && column == 1) || (line == 1 && column == x) {
				z01.PrintRune('A')
			} else if (column == 1 && line == y) || (column == x && line == y) {
				z01.PrintRune('C')
			} else if column == 1 || column == x || line == 1 || line == y {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	Raid1c(5, 3)
	z01.PrintRune('\n')
	Raid1c(5, 1)
	z01.PrintRune('\n')
	Raid1c(1, 1)
	z01.PrintRune('\n')
	Raid1c(1, 5)
}

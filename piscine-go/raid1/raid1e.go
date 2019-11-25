package main

import "github.com/01-edu/z01"

func Raid1e(x, y int) {
	for line := 1; line <= y; line++ {
		for column := 1; column <= x; column++ {
			if line == 1 && column == 1 {
				z01.PrintRune('A')
			} else if line == y && column == 1 {
				z01.PrintRune('C')
			} else if line == 1 && column == x {
				z01.PrintRune('C')
			} else if line == y && column == x {
				z01.PrintRune('A')
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
	Raid1e(5, 3)
	z01.PrintRune('\n')
	Raid1e(5, 1)
	z01.PrintRune('\n')
	Raid1e(1, 1)
	z01.PrintRune('\n')
	Raid1e(1, 5)
}

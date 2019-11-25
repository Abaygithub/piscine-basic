package main

import "github.com/01-edu/z01"

func Raid1a(x, y int) {

	for line := 1; line <= y; line++ {
		for column := 1; column <= x; column++ {
			if (line == 1 || line == y) && (column == 1 || column == x) {
				z01.PrintRune('o')
			} else if line == 1 || line == y {
				z01.PrintRune('-')
			} else if column == 1 || column == x {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

}
func main() {
	Raid1a(-10, -10)

}

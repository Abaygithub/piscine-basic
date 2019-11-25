package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x > 0 && y > 0 {
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

}

//          Ilyas code

// package student

// import "github.com/01-edu/z01"

// func Raid1a(x, y int) {
// 	if y >= 1 {
// 		for i := 1; i <= x; i++ {
// 			if i > 1 && i < x {
// 				i := '-'

// 				z01.PrintRune(i)
// 			} else {
// 				j := 'o'
// 				z01.PrintRune(j)

// 			}

// 		}
// 		z01.PrintRune('\n')

// 	}
// 	if y >= 2 {
// 		for i := 1; i <= x; i++ {
// 			if i > 1 && i < x {
// 				i := ' '

// 				z01.PrintRune(i)
// 			} else {
// 				j := '|'
// 				z01.PrintRune(j)

// 			}

// 		}
// 		z01.PrintRune('\n')
// 	}
// 	if y == 3 {
// 		for i := 1; i <= x; i++ {
// 			if i > 1 && i < x {
// 				i := '-'

// 				z01.PrintRune(i)
// 			} else {
// 				j := 'o'
// 				z01.PrintRune(j)

// 			}

// 		}
// 		z01.PrintRune('\n')
// 	}
// 	if y >= 4 {
// 		for i := 1; i <= x; i++ {
// 			if i > 1 && i < x {
// 				i := ' '

// 				z01.PrintRune(i)
// 			} else {
// 				j := '|'
// 				z01.PrintRune(j)

// 			}

// 		}
// 		z01.PrintRune('\n')
// 	}
// 	if y >= 5 {
// 		for i := 1; i <= x; i++ {
// 			if i > 1 && i < x {
// 				i := ' '

// 				z01.PrintRune(i)
// 			} else {
// 				j := '|'
// 				z01.PrintRune(j)

// 			}

// 		}
// 		z01.PrintRune('\n')
// 		for i := 1; i <= x; i++ {
// 			if i > 1 && i < x {
// 				i := '-'

// 				z01.PrintRune(i)
// 			} else {
// 				j := 'o'
// 				z01.PrintRune(j)

// 			}

// 		}
// 		z01.PrintRune('\n')
// 	}
// }

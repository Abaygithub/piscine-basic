package main

import "fmt"

func main() {
	var sudoku = [9][9]int{}
	s := 1
	for _, v := range sudoku {
		for _, vv := range v {

			fmt.Print(s)
			s++

		}

	}
}

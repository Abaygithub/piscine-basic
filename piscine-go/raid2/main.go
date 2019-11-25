package main

import (
	"fmt"
	"os"
)

func initBoard(inputArray [9]string) [9][9]int {
	board := [9][9]int{}
	for i, str := range inputArray {
		for index, item := range str {
			if item != '.' {
				board[i][index] = int(item - '0')
			}
		}
	}
	return board
}

func printBoard(board [9][9]int) {
	for _, row := range board {
		for index, item := range row {
			fmt.Print(item)
			if index != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func findFirstEmpty(board [9][9]int) [2]int {
	for i, row := range board {
		for j, item := range row {
			if item == 0 {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

func isValidRow(board [9][9]int, pos [2]int, digit int) bool {
	for _, item := range board[pos[0]] {
		if item == digit {
			return false
		}
	}
	return true
}

func isValidCol(board [9][9]int, pos [2]int, digit int) bool {
	for i := 0; i < 9; i++ {
		if board[i][pos[1]] == digit {
			return false
		}
	}
	return true
}

func isValidBox(board [9][9]int, pos [2]int, digit int) bool {

	for i := pos[0] / 3 * 3; i < pos[0]/3*3+3; i++ {
		for j := pos[1] / 3 * 3; j < pos[1]/3*3+3; j++ {
			if board[i][j] == digit {
				return false
			}
		}
	}
	return true
}

func isValidMove(board [9][9]int, pos [2]int, digit int) bool {
	return isValidRow(board, pos, digit) && isValidCol(board, pos, digit) && isValidBox(board, pos, digit)
}

func solve(board [9][9]int) bool {
	nextEmpty := findFirstEmpty(board)
	if nextEmpty[0] == -1 {
		printBoard(board)
		return true
	}

	for i := 1; i <= 9; i++ {
		if isValidMove(board, nextEmpty, i) {
			board[nextEmpty[0]][nextEmpty[1]] = i
			if solve(board) {
				return true
			}
		}
	}

	return false
}

func main() {

	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	for i := 1; i < 10; i++ {
		if len(os.Args[i]) != 9 {
			fmt.Println("Error")
			return
		}
	}

	input := [9]string{
		os.Args[1],
		os.Args[2],
		os.Args[3],
		os.Args[4],
		os.Args[5],
		os.Args[6],
		os.Args[7],
		os.Args[8],
		os.Args[9],
	}

	board := [9][9]int{}

	board = initBoard(input)

	result := solve(board)

	if !result {
		fmt.Println("Error")
	}

}

// hard level      ./raid2 "4...8...6" ".....2..." "26...1..8" ".9.7..6.." ".1.8.6.3." "..8..4.2." "8..5...64" "...2....." "5...7...2" | cat -e
// medium level    ./raid2 ".....7..." "..83...72" "....8.9.4" "8.6.3..59" "..37548.." "75..6.4.3" "6.1.2...." "27...93.." "...4....." | cat -e
// easy level      ./raid2 "8.69147.5" "...5..6.." "...76.93." "...349.5." ".5.....4." ".1.658..." ".95.76..." "..7..1..." "1.84352.7" | cat -e

package main

import (
	"fmt"
	"os"
)

func getNum(x rune) int {
	switch x {
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	default:
		return 0
	}
}

func initBoard(inputArray [9]string) [9][9]int {
	board := [9][9]int{}
	for i, str := range inputArray {
		for index, item := range str {
			if item != '.' {
				board[i][index] = getNum(item)
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

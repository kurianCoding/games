package main

import "fmt"

func displayBoard(board [][]int) {
	for row := range board {
		for col := range board[row] {
			if board[row][col] == 1 {
				fmt.Printf(" Q ")
			} else {
				fmt.Printf("[ ]")
			}
		}
		fmt.Printf("\n")
	}
}

func updateBoard(board [][]int, row, col int) {
	board[row][col] = 1
}

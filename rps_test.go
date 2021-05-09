package main

import "testing"

func TestWinner(t *testing.T) {

	if w := winner(1, 2); w != 2 {
		t.Error(w)
	}

	if w := winner(2, 3); w != 3 {
		t.Error(w)
	}

	if w := winner(1, 3); w != 1 {
		t.Error(w)
	}
}

func TestSinglePlay(t *testing.T) {
	singlePlay()
}

func TestUpdateBoard(t *testing.T) {
	board := make([][]int, 7, 7)
	for i := 0; i < 7; i++ {
		board[i] = make([]int, 7)
	}
	updateBoard(board, 0, 0)
	updateBoard(board, 1, 0)
	updateBoard(board, 1, 1)
	displayBoard(board)
}

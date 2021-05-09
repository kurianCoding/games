package main

import (
	"fmt"
	"math/rand"
	"time"
)

func winner(a, b int) int {
	// this function takes in two parameters
	// it returns an integer which says which
	// of the two won
	// 1->rock
	// 2->paper
	// 3->sciccors
	if b > a {
		temp := a
		a = b
		b = temp
	}

	if a-b > 1 {
		return b
	}

	return a
}

func printWinner(a int) string {
	players := []string{"rock", "paper", "scissors"}
	return players[a]
}

func singlePlay() {
	rand.Seed(time.Now().UnixNano())
	b := rand.Intn(4)
	var player string
	fmt.Scanf("%s", &player)

	playerMap := map[string]int{"rock": 1, "paper": 2, "scissors": 3}

	if _, ok := playerMap[player]; !ok {
		fmt.Println("invalid")
		return
	}

	a := playerMap[player]
	w := winner(a-1, b-1)
	fmt.Println(printWinner(w))
}

func main() {
	singlePlay()
}

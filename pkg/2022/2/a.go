package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// A, X(1) = Rock
// B, Y(2) = Paper
// C, Z(3) = Scissors
type Game struct {
	player1 string
	player2 string
	score   int
}

func readFile() []Game {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return make([]Game, 0)
	}
	defer file.Close()

	x := 1
	y := 2
	z := 3

	win := 6
	lose := 0
	tie := 3 

	score := map[string]int{"AX": x + tie, "AY": y + win, "AZ": z + lose, "BX": x + lose, "BY": y + tie, "BZ": z + win, "CX": x + win, "CY": y + lose, "CZ": z + tie}
	games := make([]Game, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)

		if len(split) > 0 {
			var game Game
			game.player1 = split[0]
			game.player2 = split[1]
			game.score = score[split[0]+split[1]]

			games = append(games, game)
		}
	}
	return games
}

func main() {
	games := readFile()
	total := 0

	for i := 0; i < len(games); i++ {
		total = total + games[i].score
	}

	fmt.Println(games)
	fmt.Println(total)
}

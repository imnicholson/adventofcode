package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// A = Rock
// B = Paper
// C = Scissors

// X = lose
// Y = tie
// Z = win
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

	a := 1
	b := 2
	c := 3

	win := 6
	lose := 0
	tie := 3

	score := map[string]int{"AX": lose + c, "AY": tie + a, "AZ": win + b, "BX": lose + a, "BY": tie + b, "BZ": win + c, "CX": lose + b, "CY": tie + c, "CZ": win + a}
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

	fmt.Println(total)
}

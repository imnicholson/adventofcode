package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type RuckSack struct {
	compartment1 string
	compartment2 string
	commonItem   string
	priority        int
}

func ReadFile() []RuckSack {
	priorities := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}
	rucksacks := make([]RuckSack, 0)

	file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return rucksacks
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		divider := length / 2

		var ruckSack RuckSack
		ruckSack.compartment1 = line[0:divider]
		ruckSack.compartment2 = line[divider : length]

		for _, c := range ruckSack.compartment1 {
			for _, v := range ruckSack.compartment2 {
				if c == v {
					ruckSack.commonItem = string(v)
					ruckSack.priority = priorities[ruckSack.commonItem]
					break
				}
			}
		}

		rucksacks = append(rucksacks, ruckSack)
	}
	return rucksacks
}

func main() {
	ruckSacks := ReadFile()
  total := 0

  for _, i := range ruckSacks {
    total = total + i.priority
  }
	fmt.Println(ruckSacks)
	fmt.Println(total)
}

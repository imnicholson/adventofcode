package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Group struct {
	rucksack1  string
	rucksack2  string
	rucksack3  string
	commonItem string
	priority   int
}

func ReadFile() []Group {
	priorities := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}
	groups := make([]Group, 0)

	file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return groups
	}
	defer file.Close()

	var group Group
	idx := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch idx {
		case 0:
			group.rucksack1 = line
			idx = idx + 1
		case 1:
			group.rucksack2 = line
			idx = idx + 1
		case 2:
			group.rucksack3 = line
			groups = append(groups, group)
			idx = 0

		}
	}

	for idx, it := range groups {
		for _, i := range it.rucksack1 {
			for _, j := range it.rucksack2 {
				if i == j {
					for _, k := range it.rucksack3 {

						if i == k {
							groups[idx].commonItem = string(i)
							groups[idx].priority = priorities[string(i)]
							break
						}

					}
				}
			}
		}
	}

	return groups
}

func main() {
	ruckSacks := ReadFile()
	total := 0

	for _, i := range ruckSacks {
		total = total + i.priority
	}

	fmt.Println(total)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() [][]int {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return make([][]int, 0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	allElfCalories := make([][]int, 0)
	currentElfCalories := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			allElfCalories = append(allElfCalories, currentElfCalories)
			currentElfCalories = make([]int, 0)
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("failed to convert to number: %v", err)
			}
			currentElfCalories = append(currentElfCalories, calories)
		}
	}

	return allElfCalories
}

func main() {
	allElfCalories := readFile()

	first := 0
	second := 0
	third := 0
	total := 0
	current := 0

	for index, elfCalories := range allElfCalories {
		fmt.Println(index, elfCalories)
		current = 0

		for j := 0; j < len(elfCalories); j++ {
			current = current + elfCalories[j]
		}

		if current > first {
			third = second
			second = first
			first = current
		} else if current > second {
			third = second
			second = current
		} else if current > third {
			third = current
		}
	}

	total = first + second + third
	fmt.Println(total)
}

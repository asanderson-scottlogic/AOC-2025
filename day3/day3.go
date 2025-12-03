package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var part1total = 0
	var part2total = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		part1total += getBatteriesTotalPart1(line)
		part2total += getBatteriesTotalPart2(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Printf("Part 1: %v and Part 2: %v", part1total, part2total)
}

func getBatteriesTotalPart1(line string) int {
	var batteries [2]int
	var index int

	batteries[0], index = getHighestValue(line[:len(line)-1], 0)
	batteries[1], _ = getHighestValue(line, index+1)

	return batteries[0]*10 + batteries[1]
}

func getBatteriesTotalPart2(line string) int {
	var batteries [12]int
	index := -1
	offset := 11

	for i := range 12 {
		batteries[i], index = getHighestValue(line[:len(line)-offset], index+1)
		offset--
	}

	return converToJoltage(batteries)
}

func getHighestValue(line string, index int) (int, int) {
	highest := 0
	highestIndex := 0
	var currentValue int

	for i := index; i < len(line); i++ {
		currentValue, _ = strconv.Atoi(string(line[i]))

		if currentValue > highest {
			highest = currentValue
			highestIndex = i
		}
	}

	return highest, highestIndex
}

func convertToJoltage(b [12]int) int {
	var joltage int

	for _, val := range b {
		joltage = joltage*10 + val
	}

	return joltage
}

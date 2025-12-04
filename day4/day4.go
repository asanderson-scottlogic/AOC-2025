package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const paper = "@"
const clear = "."

type position struct {
	row    int
	column int
}

func main() {

	var part1Total int

	rows := readFile("input.txt")

	for row := range rows {
		for column := 0; column < len(rows[row]); column++ {
			if checkPaperPart1(row, column, rows) {
				part1Total++
			}
		}
	}

	fmt.Printf("Part 1 total: %v", part1Total)
}

func checkPaperPart1(row, column int, rows []string) bool {
	symbol := string(rows[row][column])

	if symbol == clear {
		return false
	}

	above := row - 1
	below := row + 1
	left := column - 1
	right := column + 1
	var count int

	positions := []position{
		{above, left},
		{above, column},
		{above, right},
		{row, left},
		{row, right},
		{below, left},
		{below, column},
		{below, right},
	}

	for _, position := range positions {
		if position.row < 0 || position.row > len(rows)-1 || position.column < 0 || position.column > len(rows[row])-1 {
			continue
		}

		if string(rows[position.row][position.column]) == paper {
			count++
		}
	}

	return count < 4
}

func readFile(filename string) []string {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return lines
}

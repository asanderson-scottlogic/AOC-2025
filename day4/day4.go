package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const paper = "@"
const clear = '.'

type position struct {
	row    int
	column int
}

func main() {

	var part1Total int
	var part2Total int

	rows := readFile("input.txt")

	// part 1
	for row := range rows {
		for column := 0; column < len(rows[row]); column++ {
			if checkPaper(row, column, rows) {
				part1Total++
			}
		}
	}

	// part 2
	hits := []position{}
	previousTotal := -1

	for previousTotal != part2Total {
		previousTotal = part2Total
		for row := range rows {
			for column := 0; column < len(rows[row]); column++ {
				if checkPaper(row, column, rows) {
					part2Total++
					hits = append(hits, position{row, column})
				}
			}
		}
		for _, hit := range hits {
			rows[hit.row] = removePaperFromString(rows[hit.row], hit.column)
		}
	}

	fmt.Printf("Part 1 total: %v and Part 2 total: %v", part1Total, part2Total)
}

func checkPaper(row, column int, rows []string) bool {
	symbol := string(rows[row][column])

	if symbol != paper {
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

func removePaperFromString(input string, column int) string {
	b := []byte(input)
	b[column] = clear
	return string(b)
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

package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const splitter = "^"
const clear = "."
const start = "S"
const beam = "|"

func MainDay7(filename string) {
	lines := readFile(filename)

	fmt.Printf("Part 1 beam splits: %v", part1(lines))
}

func part1(lines []string) (splitCount int) {
	startingIndex := strings.Index(lines[0], start)
	lineLength := len(lines[0])

	// set the first beam
	lines[1] = replaceWithBeam(lines[1], startingIndex)

	for row := 2; row < len(lines); row++ {
		for column := range lineLength {
			// look above each
			above := string(lines[row-1][column])
			current := string(lines[row][column])

			if above == start {
				lines[row] = replaceWithBeam(lines[row], column)
				continue
			}

			if current == clear {
				if above == beam {
					lines[row] = replaceWithBeam(lines[row], column)
				}
				continue
			}

			if current == splitter {
				if above == beam {
					splitCount++
					// create beam left
					if column-1 >= 0 {
						lines[row] = replaceWithBeam(lines[row], column-1)
					}

					// create beam right
					if column+1 < lineLength {
						lines[row] = replaceWithBeam(lines[row], column+1)
					}
				}
			}

		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return
}

func replaceWithBeam(input string, column int) string {
	b := []byte(input)
	b[column] = '|'
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

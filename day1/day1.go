package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dial struct {
	value int
}

const left = "L"
const right = "R"

var part1total = 0
var part2total = 0

func MainDay1(input string) {
	dial := Dial{50}
	part1total = 0
	part2total = 0

	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		direction := string(line[0])
		turns, _ := strconv.Atoi(line[1:])

		dial.Turn(direction, turns)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Printf("Part 1: %v and Part 2: %v", part1total, part2total)
}

func (d *Dial) Turn(direction string, turns int) {
	switch direction {
	case left:
		for range turns {
			d.value -= 1
			if d.value < 0 {
				d.value = 99
			}
			if d.value == 0 {
				part2total++
			}
		}
	case right:
		for range turns {
			d.value += 1
			if d.value > 99 {
				d.value = 0
			}
			if d.value == 0 {
				part2total++
			}
		}
	}

	if d.value%100 == 0 {
		part1total++
	}
}

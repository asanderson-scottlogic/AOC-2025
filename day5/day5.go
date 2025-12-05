package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type inclusiveRange struct {
	lower int
	upper int
}

func MainDay5(filename string) {
	ranges, inputs := readFile(filename)
	part1(ranges, inputs)

}

func part1(ranges []string, inputs []string) {
	inclusiveRanges := convertToRangesType(ranges)
	var count int

	for _, input := range inputs {
		if checkInputInRange(input, inclusiveRanges) {
			count++
		}
	}

	fmt.Printf("Count is: %v", count)
}

func convertToRangesType(ranges []string) []inclusiveRange {
	var inclusiveRanges []inclusiveRange

	for _, r := range ranges {
		splitRange := strings.Split(r, "-")
		lower, _ := strconv.Atoi(splitRange[0])
		upper, _ := strconv.Atoi(splitRange[1])
		sr := inclusiveRange{lower, upper}
		inclusiveRanges = append(inclusiveRanges, sr)
	}

	return inclusiveRanges

}

func checkInputInRange(input string, inclusiveRanges []inclusiveRange) bool {
	for _, r := range inclusiveRanges {
		value, _ := strconv.Atoi(input)
		if value >= r.lower && value <= r.upper {
			return true
		}
	}
	return false
}

func readFile(filename string) ([]string, []string) {
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

	var ranges []string
	var inputs []string

	var section int

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			section++
			continue
		}

		if section == 0 {
			ranges = append(ranges, lines[i])
		} else {
			inputs = append(inputs, lines[i])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return ranges, inputs
}

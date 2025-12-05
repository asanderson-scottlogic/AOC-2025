package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	part2(ranges)
}

func part1(ranges []string, inputs []string) {
	inclusiveRanges := convertToRangesType(ranges)
	var count int

	for _, input := range inputs {
		if checkInputInRange(input, inclusiveRanges) {
			count++
		}
	}

	fmt.Printf("Part 1 count is: %v\n", count)
}

func part2(ranges []string) {
	inclusiveRanges := convertToRangesType(ranges)
	var count int

	sort.Slice(inclusiveRanges, func(i, j int) bool {
		return inclusiveRanges[i].lower < inclusiveRanges[j].lower
	})

	combined := combineRanges(inclusiveRanges)

	for _, c := range combined {
		count += c.upper - c.lower + 1
	}

	fmt.Printf("Part 2 count is: %v", count)

}

func combineRanges(sortedRanges []inclusiveRange) (combinedRanges []inclusiveRange) {
	for i := 0; i < len(sortedRanges); i++ {
		high := sortedRanges[0].upper

		for j := i; j < len(sortedRanges); j++ {
			low := sortedRanges[i].lower

			if sortedRanges[j].upper > high {
				high = sortedRanges[j].upper
			}

			if j+1 >= len(sortedRanges) {
				// out of bounds
				combinedRanges = append(combinedRanges, inclusiveRange{low, high})
				i = j
				break
			}

			if high < sortedRanges[j+1].lower {
				// move onto the next starting range
				combinedRanges = append(combinedRanges, inclusiveRange{low, high})
				i = j
				break
			}

		}
	}

	return combinedRanges
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

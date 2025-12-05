package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type IdRange struct {
	Start string
	End   string
}

const inputString = ""

func MainDay2() {
	var part1Sum int = 0
	var part2Sum int = 0
	var idRanges []IdRange

	ids := strings.SplitSeq(inputString, ",")

	for id := range ids {
		splitId := strings.Split(id, "-")
		entry := IdRange{splitId[0], splitId[1]}
		idRanges = append(idRanges, entry)
	}

	for _, idRange := range idRanges {
		start, _ := strconv.Atoi(idRange.Start)
		end, _ := strconv.Atoi(idRange.End)

		for id := start; id <= end; id++ {
			if checkForRepeatsPart1(id) {
				part1Sum += id
			}
			if checkForRepeatsPart2(id) {
				part2Sum += id
			}
		}
	}

	fmt.Printf("Part 1 sum: %v and Part 2 sum: %v", part1Sum, part2Sum)
}

func checkForRepeatsPart1(input int) bool {
	sequence := strconv.Itoa(input)

	if len(sequence)%2 != 0 {
		return false
	}

	midPoint := len(sequence) / 2

	firstHalf := sequence[:midPoint]
	secondHalf := sequence[midPoint:]
	return firstHalf == secondHalf
}

func checkForRepeatsPart2(input int) bool {
	sequence := strconv.Itoa(input)

	for i := 0; i < len(sequence)/2+1; i++ {
		substring := sequence[0:i]

		if (len(substring) == 0) || (len(sequence)%len(substring) != 0) {
			continue
		}

		fullString := strings.Repeat(substring, len(sequence)/len(substring))

		if fullString == sequence {
			return true
		}
	}
	return false
}

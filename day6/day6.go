package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type calculation struct {
	operation string
	length    int
}

func MainDay6(filename string) {
	lines := readFile(filename)
	lastRow := len(lines) - 1
	operationString := lines[lastRow]

	calculations := getCalculations(operationString)

	part1(calculations, lines, lastRow)
	part2(calculations, lines, lastRow)
}

func part2(calculations []calculation, lines []string, lastRow int) {
	p := 0
	var numbers []string

	runningTotal := 0
	for _, calc := range calculations {
		numbers = nil
		for i := p; i < p+calc.length; i++ {
			fullNum := ""
			for j := 0; j < lastRow; j++ {
				fullNum += removeWhitespace(string(lines[j][i]))
			}
			if fullNum != "" {
				numbers = append(numbers, fullNum)
			}
		}
		p += calc.length

		calcTotal, _ := strconv.Atoi(numbers[0])

		switch calc.operation {
		case "*":
			for i := 1; i < len(numbers); i++ {
				val, _ := strconv.Atoi(numbers[i])
				calcTotal *= val
			}
			runningTotal += calcTotal
		case "+":
			for i := 1; i < len(numbers); i++ {
				val, _ := strconv.Atoi(numbers[i])
				calcTotal += val
			}
			runningTotal += calcTotal
		}
	}

	fmt.Printf("Part 2 total is: %v", runningTotal)
}

func part1(calculations []calculation, lines []string, lastRow int) {
	p := 0
	runningTotal := 0
	for _, calc := range calculations {
		calcTotal, _ := strconv.Atoi(removeWhitespace(lines[0][p : p+calc.length]))

		switch calc.operation {
		case "*":
			for i := 1; i < lastRow; i++ {
				val, _ := strconv.Atoi(removeWhitespace(lines[i][p : p+calc.length]))
				calcTotal *= val
			}
			runningTotal += calcTotal
		case "+":
			for i := 1; i < lastRow; i++ {
				val, _ := strconv.Atoi(removeWhitespace(lines[i][p : p+calc.length]))
				calcTotal += val
			}
			runningTotal += calcTotal
		}
		p += calc.length
	}
	fmt.Printf("Part 1 total is: %v\n", runningTotal)
}

func removeWhitespace(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

func getCalculations(operationString string) (calculations []calculation) {
	for i := 0; i < len(operationString); i++ {
		currentChar := string(operationString[i])
		lenCounter := 1
		operation := currentChar

		for j := i + 1; j < len(operationString); j++ {
			nextChar := string(operationString[j])

			if (nextChar == "*") || (nextChar == "+") {
				i = j - 1
				break
			} else {
				i++
				lenCounter++
			}
		}
		calculations = append(calculations, calculation{operation: operation, length: lenCounter})
	}
	return
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

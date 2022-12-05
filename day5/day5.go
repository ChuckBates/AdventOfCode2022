package day5

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ComputePart1(input string) string {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	crateConfinguration, instructions := ParseInput(string(content))

	for _, instruction := range instructions {
		crateConfinguration = ExecuteInstructionIncrementally(instruction, crateConfinguration)
	}

	result := ""
	for _, crate := range crateConfinguration {
		result += crate[len(crate)-1]
	}

	return result
}

func ComputePart2(input string) string {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	crateConfinguration, instructions := ParseInput(string(content))

	for _, instruction := range instructions {
		crateConfinguration = ExecuteInstructionAsWhole(instruction, crateConfinguration)
	}

	result := ""
	for _, crate := range crateConfinguration {
		result += crate[len(crate)-1]
	}

	return result
}

func ParseInput(content string) ([][]string, []string) {
	emptyLineSeparator := regexp.MustCompile(`\n\s*\n`)
	parts := emptyLineSeparator.Split(string(content), -1)
	crateConfiguration := parseCrateConfiguration(parts[0])
	instructions := parseInstructions(parts[1])

	return crateConfiguration, instructions
}

func ExecuteInstructionIncrementally(instruction string, crates [][]string) [][]string {
	qty, origin, destination := parseInstruction(instruction)
	for i := 0; i < qty; i++ {
		newCrate, removed := RemoveLast(crates[origin], 1)
		crates[origin] = newCrate
		crates[destination] = append(crates[destination], removed[0])
	}
	return crates
}

func ExecuteInstructionAsWhole(instruction string, crates [][]string) [][]string {
	qty, origin, destination := parseInstruction(instruction)
	newCrate, removed := RemoveLast(crates[origin], qty)
	crates[origin] = newCrate
	crates[destination] = append(crates[destination], removed...)

	return crates
}

func RemoveLast(slice []string, qty int) ([]string, []string) {
	removedValues := []string{}
	for i := qty; i > 0; i-- {
		removedValues = append(removedValues, slice[len(slice)-i])
	}
	resultSlice := slice[:len(slice)-qty]

	return resultSlice, removedValues
}

func parseInstruction(instruction string) (int, int, int) {
	parts := strings.Split(instruction, " ")
	qty, _ := strconv.Atoi(parts[1])
	origin, _ := strconv.Atoi(parts[3])
	destination, _ := strconv.Atoi(parts[5])

	return qty, origin - 1, destination - 1
}

func parseCrateConfiguration(crateConfiguration string) [][]string {
	// "    [D]    \r\n[N] [C]    \r\n[Z] [M] [P]\r\n 1   2   3 \r"
	result := [][]string{}
	lines := strings.Split(crateConfiguration, "\r\n")
	numberOfColumns := len(trimLine(lines[0]))
	for i := 0; i < numberOfColumns; i++ {
		resultLine := []string{}
		for _, line := range lines {
			if strings.HasSuffix(line, "\r") {
				continue
			}

			trimmedLine := trimLine(line)
			valueToInsert := trimmedLine[i]
			if strings.TrimSpace(valueToInsert) == "" {
				continue
			}
			resultLine = prepend(valueToInsert, resultLine)
		}
		result = append(result, resultLine)
	}

	return result
}

func prepend(value string, array []string) []string {
	return append([]string{value}, array...)
}

func trimLine(line string) []string {
	result := []string{}

	for i := 0; i < len(line); i++ {
		result = append(result, strings.Trim(strings.Trim(line[i:i+3], "["), "]"))
		i += 3
	}

	return result
}

func parseInstructions(instructions string) []string {
	// "move 1 from 2 to 1\r\nmove 3 from 1 to 3\r\nmove 2 from 2 to 1\r\nmove 1 from 1 to 2"
	parsedInstructions := []string{}
	parsedInstructions = append(parsedInstructions, strings.Split(instructions, "\r\n")...)
	return parsedInstructions
}

package day4

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ComputePart1(input string) int {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	values := newLineSeparator.Split(string(content), -1)
	for _, value := range values {
		parts := strings.Split(strings.Trim(value, "\r"), ",")
		if oneContainsTheOther(parts[0], parts[1]) {
			sum++
		}
	}

	return sum
}

func ComputePart2(input string) int {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	values := newLineSeparator.Split(string(content), -1)
	for _, value := range values {
		parts := strings.Split(strings.Trim(value, "\r"), ",")
		if oneOverlapsTheOther(parts[0], parts[1]) {
			sum++
		}
	}

	return sum
}

func oneContainsTheOther(first string, second string) bool {
	firstParts := strings.Split(first, "-")
	firstPartOne, _ := strconv.Atoi(firstParts[0])
	firstPartTwo, _ := strconv.Atoi(firstParts[1])
	secondParts := strings.Split(second, "-")
	secondPartOne, _ := strconv.Atoi(secondParts[0])
	secondPartTwo, _ := strconv.Atoi(secondParts[1])

	return (firstPartOne >= secondPartOne && firstPartTwo <= secondPartTwo) || (secondPartOne >= firstPartOne && secondPartTwo <= firstPartTwo)
}

func oneOverlapsTheOther(first string, second string) bool {
	firstParts := strings.Split(first, "-")
	firstPartOne, _ := strconv.Atoi(firstParts[0])
	firstPartTwo, _ := strconv.Atoi(firstParts[1])
	secondParts := strings.Split(second, "-")
	secondPartOne, _ := strconv.Atoi(secondParts[0])
	secondPartTwo, _ := strconv.Atoi(secondParts[1])

	return (firstPartTwo >= secondPartOne && firstPartOne <= secondPartTwo) || (secondPartTwo >= firstPartOne && secondPartOne <= firstPartTwo)
}

package day2

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func ComputePart1(input string) int {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	values := newLineSeparator.Split(string(content), -1)
	for _, value := range values {
		choices := strings.Split(value, " ")
		total += scorePart1(choices)
	}

	return total
}

func ComputePart2(input string) int {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	values := newLineSeparator.Split(string(content), -1)
	for _, value := range values {
		choices := strings.Split(value, " ")
		total += scorePart2(choices)
	}

	return total
}

func scorePart1(choices []string) int {
	choiceScore := scoreChoice(strings.Trim(choices[1], "\r"))
	outcomeScore := scoreOutcome(choices)
	return choiceScore + outcomeScore
}

func scorePart2(choices []string) int {
	choices[1] = getChoice(choices)
	return scorePart1(choices)
}

func getChoice(choices []string) string {
	one := strings.Trim(choices[0], "\r")
	two := strings.Trim(choices[1], "\r")
	if two == "X" {
		if one == "A" {
			return "Z"
		} else if one == "B" {
			return "X"
		}
		return "Y"
	} else if two == "Y" {
		if one == "A" {
			return "X"
		} else if one == "B" {
			return "Y"
		}
		return "Z"
	}

	if one == "A" {
		return "Y"
	} else if one == "B" {
		return "Z"
	}
	return "X"
}

func scoreChoice(choice string) int {
	if choice == "X" {
		return 1
	} else if choice == "Y" {
		return 2
	} else if choice == "Z" {
		return 3
	}
	return 0
}

func scoreOutcome(choices []string) int {
	one := strings.Trim(choices[0], "\r")
	two := strings.Trim(choices[1], "\r")
	if (one == "A" && two == "X") || (one == "B" && two == "Y") || (one == "C" && two == "Z") {
		return 3
	}

	if (one == "A" && two == "Y") || (one == "C" && two == "X") || (one == "B" && two == "Z") {
		return 6
	}

	return 0
}

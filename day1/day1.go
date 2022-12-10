package day1

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

	highest := 0
	current := 0
	emptyLineSeparator := regexp.MustCompile(`\n\s*\n`)
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	parts := emptyLineSeparator.Split(string(content), -1)
	for _, part := range parts {
		values := newLineSeparator.Split(part, -1)
		for _, value := range values {
			i, err := strconv.Atoi(strings.Trim(value, "\r"))
			if err != nil {
				panic(err)
			}
			current += i
		}
		if current > highest {
			highest = current
		}
		current = 0
	}

	return highest
}

func ComputePart2(input string) int {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	highest := 0
	secondHighest := 0
	thirdHighest := 0
	current := 0
	emptyLineSeparator := regexp.MustCompile(`\n\s*\n`)
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	parts := emptyLineSeparator.Split(string(content), -1)
	for _, part := range parts {
		values := newLineSeparator.Split(part, -1)
		for _, value := range values {
			i, err := strconv.Atoi(strings.Trim(value, "\r"))
			if err != nil {
				panic(err)
			}
			current += i
		}
		if current > highest {
			thirdHighest = secondHighest
			secondHighest = highest
			highest = current
		} else if current > secondHighest {
			thirdHighest = secondHighest
			secondHighest = current
		} else if current > thirdHighest {
			thirdHighest = current
		}
		current = 0
	}

	return highest + secondHighest + thirdHighest
}

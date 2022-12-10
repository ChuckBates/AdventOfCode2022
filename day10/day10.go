package day10

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ChuckBates/AdventOfCode2022/library"
)

func ComputePart1(input string) int {
	content := library.LoadFile(input)
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	commands := newLineSeparator.Split(content, -1)
	cycleLedger := HandleCommands(commands, 1)

	return SumSignalStrengths(cycleLedger)
}

func ComputePart2(input string) {
	content := library.LoadFile(input)
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	commands := newLineSeparator.Split(content, -1)
	cycleLedger := HandleCommands(commands, 1)
	rows := EvaluateRows(cycleLedger)

	fmt.Println("Day 10 part 1:")
	for _, row := range rows {
		fmt.Println(row)
	}
}

func EvaluateRows(cycleLedger []int) []string {
	rows := []string{}

	for i := 0; i < 240; i += 40 {
		rows = append(rows, EvaluateRow(cycleLedger[i:i+40]))
	}
	return rows
}

func EvaluateRow(cycleLedger []int) string {
	row := ""
	spritePosition := 1
	for cycle := 0; cycle < len(cycleLedger); cycle++ {
		spritePosition = cycleLedger[cycle]
		if IsSpriteInView(spritePosition, cycle) {
			row += "#"
		} else {
			row += "."
		}
	}

	return row
}

func IsSpriteInView(register int, cycle int) bool {
	if cycle >= register-1 && cycle <= register+1 {
		return true
	}
	return false
}

func SumSignalStrengths(cycleLedger []int) int {
	sum := 0
	for i := 20; i < len(cycleLedger); i += 40 {
		cycle := i
		register := cycleLedger[cycle-1]
		newSum := cycle * register
		sum += newSum
	}

	return sum
}

func HandleCommands(commands []string, startRegister int) []int {
	cycleLedger := []int{}
	register := startRegister
	cycle := 1
	for _, command := range commands {
		trimmedCommand := strings.Trim(command, "\r")
		newCycle, newRegister := HandleCommand(trimmedCommand, cycle, register)
		for i := 0; i < newCycle-cycle; i++ {
			cycleLedger = append(cycleLedger, register)
		}
		register = newRegister
		cycle = newCycle
	}
	cycleLedger = append(cycleLedger, register)

	return cycleLedger
}

func HandleCommand(command string, cycle int, register int) (int, int) {
	newCycle := cycle
	newRegister := register
	commandParts := strings.Split(command, " ")

	if commandParts[0] == "noop" {
		newCycle++
	}

	if commandParts[0] == "addx" {
		newCycle += 2
		amount, _ := strconv.Atoi(commandParts[1])
		newRegister += amount
	}

	return newCycle, newRegister
}

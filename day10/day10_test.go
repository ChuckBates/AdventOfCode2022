package day10_test

import (
	"regexp"
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day10"
	"github.com/ChuckBates/AdventOfCode2022/library"
	"github.com/stretchr/testify/assert"
)

func TestComputePart1(t *testing.T) {
	assert.Equal(t, 13140, day10.ComputePart1("test.txt"))
}

func TestEvaluateRows(t *testing.T) {
	input := library.LoadFile("test.txt")
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	commands := newLineSeparator.Split(input, -1)
	cycleLedger := day10.HandleCommands(commands, 1)

	expectedRows := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}

	actualRows := day10.EvaluateRows(cycleLedger)

	assert.EqualValues(t, expectedRows, actualRows)
}

func TestEvaluateRow(t *testing.T) {
	input := library.LoadFile("test.txt")
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	commands := newLineSeparator.Split(input, -1)
	cycleLedger := day10.HandleCommands(commands, 1)
	firstFortyCycleLedger := cycleLedger[:40]

	expectedRow := "##..##..##..##..##..##..##..##..##..##.."

	actualRow := day10.EvaluateRow(firstFortyCycleLedger)

	assert.Equal(t, expectedRow, actualRow)
}

func TestIsSpriteInViewLeft(t *testing.T) {
	register := 4
	cycle := 5

	assert.True(t, day10.IsSpriteInView(register, cycle))
}

func TestIsSpriteInViewRight(t *testing.T) {
	register := 6
	cycle := 5

	assert.True(t, day10.IsSpriteInView(register, cycle))
}

func TestIsSpriteInViewCenter(t *testing.T) {
	register := 5
	cycle := 5

	assert.True(t, day10.IsSpriteInView(register, cycle))
}

func TestIsSpriteInViewOutsideLeft(t *testing.T) {
	register := 3
	cycle := 5

	assert.False(t, day10.IsSpriteInView(register, cycle))
}

func TestIsSpriteInViewOutsideRight(t *testing.T) {
	register := 7
	cycle := 5

	assert.False(t, day10.IsSpriteInView(register, cycle))
}

func TestSumSignalStrengths(t *testing.T) {
	input := library.LoadFile("test.txt")
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	commands := newLineSeparator.Split(input, -1)
	cycleLedger := day10.HandleCommands(commands, 1)
	expectedSum := 13140

	actualSum := day10.SumSignalStrengths(cycleLedger)

	assert.Equal(t, expectedSum, actualSum)
}

func TestHandleCommands(t *testing.T) {
	commands := []string{}
	commands = append(commands, "noop")
	commands = append(commands, "addx 3")
	commands = append(commands, "addx -5")

	startRegister := 1

	expectedCycleLedger := []int{}
	expectedCycleLedger = append(expectedCycleLedger, 1)
	expectedCycleLedger = append(expectedCycleLedger, 1)
	expectedCycleLedger = append(expectedCycleLedger, 1)
	expectedCycleLedger = append(expectedCycleLedger, 4)
	expectedCycleLedger = append(expectedCycleLedger, 4)
	expectedCycleLedger = append(expectedCycleLedger, -1)

	actualCycleLedger := day10.HandleCommands(commands, startRegister)

	assert.EqualValues(t, expectedCycleLedger, actualCycleLedger)
}

func TestHandleCommandNoop(t *testing.T) {
	command := "noop"
	cycle := 13
	register := 42

	expectedCycle := 14
	expectedRegister := register

	actualCycle, actualRegister := day10.HandleCommand(command, cycle, register)
	assert.Equal(t, expectedCycle, actualCycle)
	assert.Equal(t, expectedRegister, actualRegister)
}

func TestHandleCommandAdd(t *testing.T) {
	command := "addx 3"
	cycle := 2
	register := 1

	expectedCycle := 4
	expectedRegister := 4

	actualCycle, actualRegister := day10.HandleCommand(command, cycle, register)
	assert.Equal(t, expectedCycle, actualCycle)
	assert.Equal(t, expectedRegister, actualRegister)
}

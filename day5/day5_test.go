package day5_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day5"
	"github.com/ChuckBates/AdventOfCode2022/library"
	"github.com/stretchr/testify/assert"
)

func TestComputePart1(t *testing.T) {
	if day5.ComputePart1("test.txt") != "CMZ" {
		t.Fatal("Wrong value")
	}
}

func TestComputePart2(t *testing.T) {
	if day5.ComputePart2("test.txt") != "MCD" {
		t.Fatal("Wrong value")
	}
}

func TestParseInput(t *testing.T) {
	expectedCrateConfiguration := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	expectedInstructions := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	input := library.LoadFile("test.txt")

	actualCrateConfiguration, actualInstructions := day5.ParseInput(input)

	assert.Equal(t, expectedCrateConfiguration, actualCrateConfiguration)
	assert.Equal(t, expectedInstructions, actualInstructions)
}

func TestExecuteInstructionIncrementally(t *testing.T) {
	expectedStartCrates := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	expectedInstructions := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	expectedCratesEvolution := [][][]string{
		{
			{"Z", "N", "D"},
			{"M", "C"},
			{"P"},
		},
		{
			{},
			{"M", "C"},
			{"P", "D", "N", "Z"},
		},
		{
			{"C", "M"},
			{},
			{"P", "D", "N", "Z"},
		},
		{
			{"C"},
			{"M"},
			{"P", "D", "N", "Z"},
		},
	}

	input := library.LoadFile("test.txt")

	actualCrateConfiguration, actualInstructions := day5.ParseInput(input)

	assert.Equal(t, expectedStartCrates, actualCrateConfiguration)
	assert.Equal(t, expectedInstructions, actualInstructions)

	previousCrateConfig := actualCrateConfiguration
	for index, instruction := range actualInstructions {
		previousCrateConfig = day5.ExecuteInstructionIncrementally(instruction, previousCrateConfig)
		assert.Equal(t, expectedCratesEvolution[index], previousCrateConfig)
	}
}

func TestExecuteInstructionAsWhole(t *testing.T) {
	expectedStartCrates := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	expectedInstructions := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	expectedCratesEvolution := [][][]string{
		{
			{"Z", "N", "D"},
			{"M", "C"},
			{"P"},
		},
		{
			{},
			{"M", "C"},
			{"P", "Z", "N", "D"},
		},
		{
			{"M", "C"},
			{},
			{"P", "Z", "N", "D"},
		},
		{
			{"M"},
			{"C"},
			{"P", "Z", "N", "D"},
		},
	}

	input := library.LoadFile("test.txt")

	actualCrateConfiguration, actualInstructions := day5.ParseInput(input)

	assert.Equal(t, expectedStartCrates, actualCrateConfiguration)
	assert.Equal(t, expectedInstructions, actualInstructions)

	previousCrateConfig := actualCrateConfiguration
	for index, instruction := range actualInstructions {
		previousCrateConfig = day5.ExecuteInstructionAsWhole(instruction, previousCrateConfig)
		assert.Equal(t, expectedCratesEvolution[index], previousCrateConfig)
	}
}

func TestRemoveLastSingle(t *testing.T) {
	inputSlice := []string{"A", "B", "C", "D"}
	expectedSlice := []string{"A", "B", "C"}
	expectedValues := []string{"D"}

	actualSlice, removedValues := day5.RemoveLast(inputSlice, 1)

	assert.Equal(t, expectedSlice, actualSlice)
	assert.Equal(t, expectedValues, removedValues)
}

func TestRemoveLastGroup(t *testing.T) {
	inputSlice := []string{"A", "B", "C", "D"}
	expectedSlice := []string{"A"}
	expectedValues := []string{"B", "C", "D"}

	actualSlice, removedValues := day5.RemoveLast(inputSlice, 3)

	assert.Equal(t, expectedSlice, actualSlice)
	assert.Equal(t, expectedValues, removedValues)
}

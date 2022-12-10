package day9

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ChuckBates/AdventOfCode2022/library"
)

type Instruction struct {
	Direction string
	Amount    int
}

type Knot struct {
	CurrentPosition [2]int
	AllPositions    [][2]int
}

func ComputePart2(input string) int {
	content := library.LoadFile(input)
	instructions := ParseInstructions(content)

	knots := BuildStartingKnots([2]int{15, 11}, 10)
	for _, instruction := range instructions {
		knots = ApplyInstruction(instruction, knots)
	}

	return len(knots[len(knots)-1].AllPositions)
}

func ComputePart1(input string) int {
	content := library.LoadFile(input)
	instructions := ParseInstructions(content)

	knots := BuildStartingKnots([2]int{4, 0}, 2)
	for _, instruction := range instructions {
		knots = ApplyInstruction(instruction, knots)
	}

	return len(knots[len(knots)-1].AllPositions)
}

func BuildStartingKnots(startPosition [2]int, numberOfKnots int) []Knot {
	knots := []Knot{}
	for i := 0; i < numberOfKnots; i++ {
		knotPositions := [][2]int{startPosition}
		knots = append(knots, Knot{startPosition, knotPositions})
	}

	return knots
}

func HasPosition(positions [][2]int, newPosition [2]int) bool {
	for _, position := range positions {
		if position[0] == newPosition[0] && position[1] == newPosition[1] {
			return true
		}
	}

	return false
}

func UpdateKnots(knots []Knot, xDelta int, yDelta int) []Knot {
	newKnots := knots

	newKnotOnePosition := [2]int{newKnots[0].CurrentPosition[0] + xDelta, newKnots[0].CurrentPosition[1] + yDelta}
	newKnots[0].CurrentPosition = newKnotOnePosition
	if !HasPosition(newKnots[0].AllPositions, newKnotOnePosition) {
		newKnots[0].AllPositions = append(newKnots[0].AllPositions, newKnotOnePosition)
	}

	for j := 1; j < len(newKnots); j++ {
		if KnotNeedsToMove(newKnots[j-1].CurrentPosition, newKnots[j].CurrentPosition) {
			newKnotPosition := MoveKnotTwoToKnotOne(newKnots[j-1].CurrentPosition, newKnots[j].CurrentPosition)
			newKnots[j].CurrentPosition = newKnotPosition
			if !HasPosition(newKnots[j].AllPositions, newKnotPosition) {
				newKnots[j].AllPositions = append(newKnots[j].AllPositions, newKnotPosition)
			}
		}
	}

	return newKnots
}

func ApplyInstruction(instruction Instruction, knots []Knot) []Knot {
	updatedKnots := knots

	if instruction.Direction == "R" {
		for i := 1; i <= instruction.Amount; i++ {
			updatedKnots = UpdateKnots(updatedKnots, 0, 1)
		}
	} else if instruction.Direction == "L" {
		for i := 1; i <= instruction.Amount; i++ {
			updatedKnots = UpdateKnots(updatedKnots, 0, -1)
		}
	} else if instruction.Direction == "U" {
		for i := 1; i <= instruction.Amount; i++ {
			updatedKnots = UpdateKnots(updatedKnots, -1, 0)
		}
	} else if instruction.Direction == "D" {
		for i := 1; i <= instruction.Amount; i++ {
			updatedKnots = UpdateKnots(updatedKnots, 1, 0)
		}
	}

	return updatedKnots
}

func ParseInstructions(input string) []Instruction {
	instructions := []Instruction{}

	newLineSeparator := regexp.MustCompile(`\n\s*`)
	lines := newLineSeparator.Split(input, -1)
	for _, line := range lines {
		lineParts := strings.Split(strings.Trim(line, "\r"), " ")
		amount, _ := strconv.Atoi(lineParts[1])
		instructions = append(instructions, Instruction{lineParts[0], amount})
	}

	return instructions
}

func KnotNeedsToMove(knotOnePosition [2]int, knotTwoPosition [2]int) bool {
	xDelta := math.Abs(float64(knotOnePosition[0] - knotTwoPosition[0]))
	yDelta := math.Abs(float64(knotOnePosition[1] - knotTwoPosition[1]))
	return xDelta > 1 || yDelta > 1
}

func MoveKnotTwoToKnotOne(knotOnePosition [2]int, knotTwoPosition [2]int) [2]int {
	newKnotTwoPosition := [2]int{knotTwoPosition[0], knotTwoPosition[1]}

	moveKnotRight := knotOnePosition[1]-knotTwoPosition[1] > 1
	moveKnotLeft := knotTwoPosition[1]-knotOnePosition[1] > 1
	moveKnotDown := knotOnePosition[0]-knotTwoPosition[0] > 1
	moveKnotUp := knotTwoPosition[0]-knotOnePosition[0] > 1

	if moveKnotUp && moveKnotLeft {
		newKnotTwoPosition[0] = knotOnePosition[0] + 1
		newKnotTwoPosition[1] = knotOnePosition[1] + 1
		return newKnotTwoPosition
	}

	if moveKnotDown && moveKnotLeft {
		newKnotTwoPosition[0] = knotOnePosition[0] - 1
		newKnotTwoPosition[1] = knotOnePosition[1] + 1
		return newKnotTwoPosition
	}

	if moveKnotUp && moveKnotRight {
		newKnotTwoPosition[0] = knotOnePosition[0] + 1
		newKnotTwoPosition[1] = knotOnePosition[1] - 1
		return newKnotTwoPosition
	}

	if moveKnotDown && moveKnotRight {
		newKnotTwoPosition[0] = knotOnePosition[0] - 1
		newKnotTwoPosition[1] = knotOnePosition[1] - 1
		return newKnotTwoPosition
	}

	if moveKnotDown {
		newKnotTwoPosition[0] = knotOnePosition[0] - 1
		newKnotTwoPosition[1] = knotOnePosition[1]
		return newKnotTwoPosition
	}

	if moveKnotUp {
		newKnotTwoPosition[0] = knotOnePosition[0] + 1
		newKnotTwoPosition[1] = knotOnePosition[1]
		return newKnotTwoPosition
	}

	if moveKnotLeft {
		newKnotTwoPosition[0] = knotOnePosition[0]
		newKnotTwoPosition[1] = knotOnePosition[1] + 1
		return newKnotTwoPosition
	}

	if moveKnotRight {
		newKnotTwoPosition[0] = knotOnePosition[0]
		newKnotTwoPosition[1] = knotOnePosition[1] - 1
		return newKnotTwoPosition
	}

	return newKnotTwoPosition
}

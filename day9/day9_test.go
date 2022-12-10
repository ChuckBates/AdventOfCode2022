package day9_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day9"
	"github.com/ChuckBates/AdventOfCode2022/library"
	"github.com/stretchr/testify/assert"
)

func TestComputePart1(t *testing.T) {
	assert.Equal(t, 13, day9.ComputePart1("test.txt"))
}

func TestComputePart2(t *testing.T) {
	assert.Equal(t, 36, day9.ComputePart2("test2.txt"))
}

func TestNeedsToMoveWhenOnTopOfEachOther(t *testing.T) {
	FirstPosition, SecondPosition := [2]int{2, 2}, [2]int{2, 2}

	assert.False(t, day9.KnotNeedsToMove(FirstPosition, SecondPosition))
}

func TestNeedsToMoveWhenSideOfEachOther(t *testing.T) {
	FirstPosition, SecondPosition := [2]int{2, 2}, [2]int{2, 1}

	assert.False(t, day9.KnotNeedsToMove(FirstPosition, SecondPosition))
}

func TestNeedsToMoveWhenDiagonalOfEachOther(t *testing.T) {
	FirstPosition, SecondPosition := [2]int{2, 2}, [2]int{1, 1}

	assert.False(t, day9.KnotNeedsToMove(FirstPosition, SecondPosition))
}

func TestNeedsToMoveWhenTwoStepsToSideOfEachOther(t *testing.T) {
	FirstPosition, SecondPosition := [2]int{2, 2}, [2]int{0, 2}

	assert.True(t, day9.KnotNeedsToMove(FirstPosition, SecondPosition))
}

func TestNeedsToMoveWhenTwoStepsDiagonallyOfEachOther(t *testing.T) {
	FirstPosition, SecondPosition := [2]int{2, 2}, [2]int{0, 0}

	assert.True(t, day9.KnotNeedsToMove(FirstPosition, SecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsLeft(t *testing.T) {
	FirstPostition := [2]int{4, 2}
	startingSecondPosition := [2]int{4, 0}
	expectedSecondPosition := [2]int{4, 1}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsRight(t *testing.T) {
	FirstPostition := [2]int{0, 1}
	startingSecondPosition := [2]int{0, 3}
	expectedSecondPosition := [2]int{0, 2}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsUp(t *testing.T) {
	FirstPostition := [2]int{2, 2}
	startingSecondPosition := [2]int{0, 2}
	expectedSecondPosition := [2]int{1, 2}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsDown(t *testing.T) {
	FirstPostition := [2]int{2, 2}
	startingSecondPosition := [2]int{4, 2}
	expectedSecondPosition := [2]int{3, 2}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsLeftDown(t *testing.T) {
	FirstPostition := [2]int{2, 4}
	startingSecondPosition := [2]int{4, 3}
	expectedSecondPosition := [2]int{3, 4}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsLeftUp(t *testing.T) {
	FirstPostition := [2]int{2, 4}
	startingSecondPosition := [2]int{0, 3}
	expectedSecondPosition := [2]int{1, 4}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsRightUp(t *testing.T) {
	FirstPostition := [2]int{2, 2}
	startingSecondPosition := [2]int{1, 4}
	expectedSecondPosition := [2]int{2, 3}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsRightDown(t *testing.T) {
	FirstPostition := [2]int{2, 2}
	startingSecondPosition := [2]int{3, 4}
	expectedSecondPosition := [2]int{2, 3}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestMoveSecondToFirstWhenSecondIsDiagonal(t *testing.T) {
	FirstPostition := [2]int{2, 2}
	startingSecondPosition := [2]int{4, 0}
	expectedSecondPosition := [2]int{3, 1}

	assert.Equal(t, expectedSecondPosition, day9.MoveKnotTwoToKnotOne(FirstPostition, startingSecondPosition))
}

func TestParseInstructions(t *testing.T) {
	expectedInstructions := []day9.Instruction{}
	expectedInstructions = append(expectedInstructions, day9.Instruction{"R", 4})
	expectedInstructions = append(expectedInstructions, day9.Instruction{"U", 4})
	expectedInstructions = append(expectedInstructions, day9.Instruction{"L", 3})
	expectedInstructions = append(expectedInstructions, day9.Instruction{"D", 1})
	expectedInstructions = append(expectedInstructions, day9.Instruction{"R", 4})
	expectedInstructions = append(expectedInstructions, day9.Instruction{"D", 1})
	expectedInstructions = append(expectedInstructions, day9.Instruction{"L", 5})
	expectedInstructions = append(expectedInstructions, day9.Instruction{"R", 2})

	assert.EqualValues(t, expectedInstructions, day9.ParseInstructions(library.LoadFile("test.txt")))
}

func TestHasPosition(t *testing.T) {
	positions := [][2]int{}
	position := [2]int{1, 4}
	positions = append(positions, position)

	assert.True(t, day9.HasPosition(positions, position))
}

func TestBuildStartingKnots(t *testing.T) {
	startPosition := [2]int{5, 5}
	numberOfKnots := 4

	expectedKnots := []day9.Knot{}
	expectedKnots = append(expectedKnots, day9.Knot{startPosition, [][2]int{startPosition}})
	expectedKnots = append(expectedKnots, day9.Knot{startPosition, [][2]int{startPosition}})
	expectedKnots = append(expectedKnots, day9.Knot{startPosition, [][2]int{startPosition}})
	expectedKnots = append(expectedKnots, day9.Knot{startPosition, [][2]int{startPosition}})

	actualKnots := day9.BuildStartingKnots(startPosition, numberOfKnots)

	assert.Equal(t, numberOfKnots, len(actualKnots))
	assert.EqualValues(t, expectedKnots, actualKnots)
	assert.EqualValues(t, expectedKnots[0].CurrentPosition, actualKnots[0].CurrentPosition)
	assert.EqualValues(t, expectedKnots[0].AllPositions, actualKnots[0].AllPositions)
	assert.EqualValues(t, expectedKnots[1].CurrentPosition, actualKnots[1].CurrentPosition)
	assert.EqualValues(t, expectedKnots[1].AllPositions, actualKnots[1].AllPositions)
	assert.EqualValues(t, expectedKnots[2].CurrentPosition, actualKnots[2].CurrentPosition)
	assert.EqualValues(t, expectedKnots[2].AllPositions, actualKnots[2].AllPositions)
	assert.EqualValues(t, expectedKnots[3].CurrentPosition, actualKnots[3].CurrentPosition)
	assert.EqualValues(t, expectedKnots[3].AllPositions, actualKnots[3].AllPositions)
}

func TestUpdateKnotsRightSequence(t *testing.T) {
	startPosition := [2]int{0, 0}
	startKnots := day9.BuildStartingKnots(startPosition, 4)

	expectedKnots := []day9.Knot{}

	expectedFinalKnotOnePosition := [2]int{0, 4}
	expectedFinalKnotOnePositions := [][2]int{}
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, startPosition)
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, [2]int{0, 1})
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, [2]int{0, 2})
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, [2]int{0, 3})
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, expectedFinalKnotOnePosition)
	expectedFinalKnotOne := day9.Knot{expectedFinalKnotOnePosition, expectedFinalKnotOnePositions}

	expectedFinalKnotTwoPosition := [2]int{0, 3}
	expectedFinalKnotTwoPositions := [][2]int{}
	expectedFinalKnotTwoPositions = append(expectedFinalKnotTwoPositions, startPosition)
	expectedFinalKnotTwoPositions = append(expectedFinalKnotTwoPositions, [2]int{0, 1})
	expectedFinalKnotTwoPositions = append(expectedFinalKnotTwoPositions, [2]int{0, 2})
	expectedFinalKnotTwoPositions = append(expectedFinalKnotTwoPositions, expectedFinalKnotTwoPosition)
	expectedFinalKnotTwo := day9.Knot{expectedFinalKnotTwoPosition, expectedFinalKnotTwoPositions}

	expectedFinalKnotThreePosition := [2]int{0, 2}
	expectedFinalKnotThreePositions := [][2]int{}
	expectedFinalKnotThreePositions = append(expectedFinalKnotThreePositions, startPosition)
	expectedFinalKnotThreePositions = append(expectedFinalKnotThreePositions, [2]int{0, 1})
	expectedFinalKnotThreePositions = append(expectedFinalKnotThreePositions, expectedFinalKnotThreePosition)
	expectedFinalKnotThree := day9.Knot{expectedFinalKnotThreePosition, expectedFinalKnotThreePositions}

	expectedFinalKnotFourPosition := [2]int{0, 1}
	expectedFinalKnotFourPositions := [][2]int{}
	expectedFinalKnotFourPositions = append(expectedFinalKnotFourPositions, startPosition)
	expectedFinalKnotFourPositions = append(expectedFinalKnotFourPositions, expectedFinalKnotFourPosition)
	expectedFinalKnotFour := day9.Knot{expectedFinalKnotFourPosition, expectedFinalKnotFourPositions}

	expectedKnots = append(expectedKnots, expectedFinalKnotOne)
	expectedKnots = append(expectedKnots, expectedFinalKnotTwo)
	expectedKnots = append(expectedKnots, expectedFinalKnotThree)
	expectedKnots = append(expectedKnots, expectedFinalKnotFour)

	actualKnots := startKnots
	for i := 0; i < 4; i++ {
		actualKnots = day9.UpdateKnots(actualKnots, 0, 1)
	}

	assert.EqualValues(t, expectedKnots, actualKnots)
}

func TestUpdateKnotsDownSequence(t *testing.T) {
	startPosition := [2]int{0, 0}
	startKnots := day9.BuildStartingKnots(startPosition, 4)

	expectedKnots := []day9.Knot{}

	expectedFinalKnotOnePosition := [2]int{4, 0}
	expectedFinalKnotOnePositions := [][2]int{}
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, startPosition)
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, [2]int{1, 0})
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, [2]int{2, 0})
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, [2]int{3, 0})
	expectedFinalKnotOnePositions = append(expectedFinalKnotOnePositions, expectedFinalKnotOnePosition)
	expectedFinalKnotOne := day9.Knot{expectedFinalKnotOnePosition, expectedFinalKnotOnePositions}

	expectedFinalKnotTwoPosition := [2]int{3, 0}
	expectedFinalKnotTwoPositions := [][2]int{}
	expectedFinalKnotTwoPositions = append(expectedFinalKnotTwoPositions, startPosition)
	expectedFinalKnotTwoPositions = append(expectedFinalKnotTwoPositions, [2]int{1, 0})
	expectedFinalKnotTwoPositions = append(expectedFinalKnotTwoPositions, [2]int{2, 0})
	expectedFinalKnotTwoPositions = append(expectedFinalKnotTwoPositions, expectedFinalKnotTwoPosition)
	expectedFinalKnotTwo := day9.Knot{expectedFinalKnotTwoPosition, expectedFinalKnotTwoPositions}

	expectedFinalKnotThreePosition := [2]int{2, 0}
	expectedFinalKnotThreePositions := [][2]int{}
	expectedFinalKnotThreePositions = append(expectedFinalKnotThreePositions, startPosition)
	expectedFinalKnotThreePositions = append(expectedFinalKnotThreePositions, [2]int{1, 0})
	expectedFinalKnotThreePositions = append(expectedFinalKnotThreePositions, expectedFinalKnotThreePosition)
	expectedFinalKnotThree := day9.Knot{expectedFinalKnotThreePosition, expectedFinalKnotThreePositions}

	expectedFinalKnotFourPosition := [2]int{1, 0}
	expectedFinalKnotFourPositions := [][2]int{}
	expectedFinalKnotFourPositions = append(expectedFinalKnotFourPositions, startPosition)
	expectedFinalKnotFourPositions = append(expectedFinalKnotFourPositions, expectedFinalKnotFourPosition)
	expectedFinalKnotFour := day9.Knot{expectedFinalKnotFourPosition, expectedFinalKnotFourPositions}

	expectedKnots = append(expectedKnots, expectedFinalKnotOne)
	expectedKnots = append(expectedKnots, expectedFinalKnotTwo)
	expectedKnots = append(expectedKnots, expectedFinalKnotThree)
	expectedKnots = append(expectedKnots, expectedFinalKnotFour)

	actualKnots := startKnots
	for i := 0; i < 4; i++ {
		actualKnots = day9.UpdateKnots(actualKnots, 1, 0)
	}

	assert.EqualValues(t, expectedKnots, actualKnots)
}

func TestUpdateKnotsCustomSequence(t *testing.T) {
	startingKnots := []day9.Knot{}
	startingKnots = append(startingKnots, day9.Knot{[2]int{1, 5}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{2, 5}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{3, 4}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{3, 3}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{3, 2}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{3, 1}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{4, 0}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{4, 0}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{4, 0}, [][2]int{}})
	startingKnots = append(startingKnots, day9.Knot{[2]int{4, 0}, [][2]int{}})

	expectedKnots := []day9.Knot{}
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{0, 5}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{1, 5}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{2, 5}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{2, 4}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{2, 3}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{2, 2}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{3, 1}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{4, 0}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{4, 0}, [][2]int{}})
	expectedKnots = append(expectedKnots, day9.Knot{[2]int{4, 0}, [][2]int{}})

	actualKnots := day9.UpdateKnots(startingKnots, -1, 0)

	for i := 0; i < len(actualKnots); i++ {
		assert.EqualValues(t, expectedKnots[i].CurrentPosition, actualKnots[i].CurrentPosition)
	}
}

package day8_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day8"
	"github.com/stretchr/testify/assert"
)

func TestComputePart1(t *testing.T) {
	assert.Equal(t, 21, day8.ComputePart1("test.txt"))
}

func TestComputePart2(t *testing.T) {
	assert.Equal(t, 8, day8.ComputePart2("test.txt"))
}

func TestGetBestScenicScore(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	assert.Equal(t, 8, day8.GetBestScenicScore(grid))
}

func TestGetScenicScore(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	positionOne := []int{1, 2}
	assert.Equal(t, 4, day8.GetScenicScore(grid, positionOne))

	positionTwo := []int{3, 2}
	assert.Equal(t, 8, day8.GetScenicScore(grid, positionTwo))
}

func TestCountLeftVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{3, 2}

	_, count := day8.LeftVisible(grid, position)
	assert.Equal(t, 2, count)
}

func TestIsLeftVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{3, 2}

	isLeftVisible, _ := day8.LeftVisible(grid, position)
	assert.True(t, isLeftVisible)
}

func TestIsLeftNotVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{2, 2}

	isLeftVisible, _ := day8.LeftVisible(grid, position)
	assert.False(t, isLeftVisible)
}

func TestIsRightVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{1, 2}

	isRightVisible, _ := day8.RightVisible(grid, position)
	assert.True(t, isRightVisible)
}

func TestIsRightNotVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{2, 2}

	isRightVisible, _ := day8.RightVisible(grid, position)
	assert.False(t, isRightVisible)
}

func TestIsUpVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 2, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{2, 3}

	isUpVisible, _ := day8.UpVisible(grid, position)
	assert.True(t, isUpVisible)
}

func TestIsUpNotVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{2, 3}

	isUpVisible, _ := day8.UpVisible(grid, position)
	assert.False(t, isUpVisible)
}

func TestIsDownVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 2, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 4, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{1, 2}

	isDownVisible, _ := day8.DownVisible(grid, position)
	assert.True(t, isDownVisible)
}

func TestIsDownNotVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})

	position := []int{2, 2}

	isDownVisible, _ := day8.DownVisible(grid, position)
	assert.False(t, isDownVisible)
}

func TestCountVisible(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{3, 0, 3, 7, 3})
	grid = append(grid, []int{2, 5, 5, 1, 2})
	grid = append(grid, []int{6, 5, 3, 3, 2})
	grid = append(grid, []int{3, 3, 5, 4, 9})
	grid = append(grid, []int{3, 5, 3, 9, 0})
	grid = append(grid, []int{3, 5, 2, 9, 0})

	assert.Equal(t, 26, day8.CountVisible(grid))
}

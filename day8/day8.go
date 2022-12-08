package day8

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/ChuckBates/AdventOfCode2022/library"
)

func ComputePart1(input string) int {

	content := library.LoadFile(input)

	return CountVisible(BuildGrid(content))
}

func ComputePart2(input string) int {

	content := library.LoadFile(input)

	return GetBestScenicScore(BuildGrid(content))
}

func BuildGrid(content string) [][]int {
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	lines := newLineSeparator.Split(content, -1)

	grid := [][]int{}
	for _, line := range lines {
		trimmedLine := strings.Trim(line, "\r")
		row := []int{}
		for y := 0; y < len(trimmedLine); y++ {
			value, _ := strconv.Atoi(string(trimmedLine[y]))
			row = append(row, value)
		}
		grid = append(grid, row)
	}

	return grid
}

func GetBestScenicScore(grid [][]int) int {
	verticalRange := len(grid) - 2
	horizontalRange := len(grid[0]) - 2

	bestScore := 0
	for y := 1; y <= verticalRange; y++ {
		for x := 1; x <= horizontalRange; x++ {
			score := GetScenicScore(grid, []int{y, x})
			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func GetScenicScore(grid [][]int, position []int) int {
	_, leftCount := LeftVisible(grid, position)
	_, rightCount := RightVisible(grid, position)
	_, upCount := UpVisible(grid, position)
	_, downCount := DownVisible(grid, position)

	return leftCount * rightCount * upCount * downCount
}

func CountVisible(grid [][]int) int {
	verticalEdges := len(grid) * 2
	horizonalEdges := (len(grid[0]) - 2) * 2

	verticalRange := len(grid) - 2
	horizontalRange := len(grid[0]) - 2

	visible := 0
	for y := 1; y <= verticalRange; y++ {
		for x := 1; x <= horizontalRange; x++ {
			leftVisible, _ := LeftVisible(grid, []int{y, x})
			rightVisible, _ := RightVisible(grid, []int{y, x})
			upVisible, _ := UpVisible(grid, []int{y, x})
			downVisible, _ := DownVisible(grid, []int{y, x})
			if leftVisible || rightVisible || upVisible || downVisible {
				visible += 1
			}
		}
	}

	return verticalEdges + horizonalEdges + visible
}

func LeftVisible(grid [][]int, position []int) (bool, int) {
	max := grid[position[0]][position[1]]
	count := 0

	for i := position[1] - 1; i >= 0; i-- {
		adjacent := grid[position[0]][i]
		count++
		if adjacent >= max {
			return false, count
		}
	}
	return true, count
}

func RightVisible(grid [][]int, position []int) (bool, int) {
	max := grid[position[0]][position[1]]
	count := 0

	for i := position[1] + 1; i < len(grid[0]); i++ {
		adjacent := grid[position[0]][i]
		count++
		if adjacent >= max {
			return false, count
		}
	}
	return true, count
}

func UpVisible(grid [][]int, position []int) (bool, int) {
	max := grid[position[0]][position[1]]
	count := 0

	for i := position[0] - 1; i >= 0; i-- {
		adjacent := grid[i][position[1]]
		count++
		if adjacent >= max {
			return false, count
		}
	}
	return true, count
}

func DownVisible(grid [][]int, position []int) (bool, int) {
	max := grid[position[0]][position[1]]
	count := 0

	for i := position[0] + 1; i < len(grid[1]); i++ {
		adjacent := grid[i][position[1]]
		count++
		if adjacent >= max {
			return false, count
		}
	}
	return true, count
}

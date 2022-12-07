package day7

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/ChuckBates/AdventOfCode2022/library"
)

var top = &Directory{"/", []int{}, []*Directory{}, &Directory{}}
var current = top

func ComputePart1(input string) int {
	top = &Directory{"/", []int{}, []*Directory{}, &Directory{}}
	content := library.LoadFile(input)

	newLineSeparator := regexp.MustCompile(`\n\s*`)
	lines := newLineSeparator.Split(content, -1)

	for _, line := range lines {
		trimmedLine := strings.Trim(line, "\r")
		if trimmedLine == "$ cd /" {
			current = top
			continue
		}

		HandleLine(trimmedLine)
	}

	return sumDirectoriesUnderHundredThousand(top)
}

func ComputePart2(input string) int {
	top = &Directory{"/", []int{}, []*Directory{}, &Directory{}}
	content := library.LoadFile(input)

	newLineSeparator := regexp.MustCompile(`\n\s*`)
	lines := newLineSeparator.Split(content, -1)

	for _, line := range lines {
		trimmedLine := strings.Trim(line, "\r")
		if trimmedLine == "$ cd /" {
			current = top
			continue
		}

		HandleLine(trimmedLine)
	}

	totalSize := getDirectorySize(top)
	unusedSpace := 70000000 - totalSize
	totalSpaceNeeded := 30000000
	deltaSpaceNeeded := totalSpaceNeeded - unusedSpace

	return getSmallestDirectoryOverDelta(deltaSpaceNeeded, 70000000, top)
}

type Directory struct {
	Name           string
	FileSizes      []int
	SubDirectories []*Directory
	Parent         *Directory
}

func HandleLine(line string) {
	if strings.HasPrefix(line, "$") {
		HandleCommand(line)
		return
	}

	if strings.HasPrefix(line, "dir") {
		HandleNewDirectory(line)
		return
	}

	HandleFile(line)
}

func HandleCommand(line string) {
	parts := strings.Split(line, " ")
	if parts[1] == "cd" {
		if parts[2] == ".." {
			current = MoveUpDirectory()
		} else {
			containsSubDirectory, subDirectory := containsSubDirectory(parts[2])
			if containsSubDirectory {
				current = subDirectory
			}
		}
	}
}

func HandleNewDirectory(line string) {
	parts := strings.Split(line, " ")
	current.SubDirectories = append(current.SubDirectories, &Directory{parts[1], []int{}, []*Directory{}, current})
}

func HandleFile(line string) {
	parts := strings.Split(line, " ")

	size, _ := strconv.Atoi(parts[0])
	current.FileSizes = append(current.FileSizes, size)
}

func containsSubDirectory(subDirectoryName string) (bool, *Directory) {
	for _, subDirectory := range current.SubDirectories {
		if subDirectory.Name == subDirectoryName {
			return true, subDirectory
		}
	}

	return false, &Directory{}
}

func MoveUpDirectory() *Directory {
	return current.Parent
}

func getSmallestDirectoryOverDelta(delta int, currentSmallestSize int, directory *Directory) int {
	smallestSize := currentSmallestSize
	currentDirectorySize := getDirectorySize(directory)
	if currentDirectorySize >= delta && currentDirectorySize < smallestSize {
		smallestSize = currentDirectorySize
	}

	for _, sub := range directory.SubDirectories {
		smallestSize = getSmallestDirectoryOverDelta(delta, smallestSize, sub)
	}

	return smallestSize
}

func sumDirectoriesUnderHundredThousand(directory *Directory) int {
	sum := 0
	size := getDirectorySize(directory)
	if size <= 100000 {
		sum += size
	}

	for _, sub := range directory.SubDirectories {
		sum += sumDirectoriesUnderHundredThousand(sub)
	}

	return sum
}

func getDirectorySize(directory *Directory) int {
	size := 0
	for _, s := range directory.FileSizes {
		size += s
	}

	for _, sub := range directory.SubDirectories {
		size += getDirectorySize(sub)
	}

	return size
}

//Test helper functions

func GetCurrentDirectory() *Directory {
	return current
}

func SetCurrentDirectory(directory *Directory) {
	current = directory
}

func GetTopDirectory() *Directory {
	return top
}

func SetTopDirectory(directory *Directory) {
	top = directory
}

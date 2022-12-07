package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ChuckBates/AdventOfCode2022/library"
)

var top = &Directory{"/", []int{}, []*Directory{}, ""}
var current = top

func ComputePart1(input string) int {
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

	print(top)
	return sumDirectoriesUnderHundredThousand(top)
}

type Directory struct {
	name           string
	fileSizes      []int
	subDirectories []*Directory
	parentName     string
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
			current = moveUpDirectory(current.parentName)
		} else {
			containsSubDirectory, subDirectory := containsSubDirectory(parts[2])
			if containsSubDirectory {
				current = subDirectory
			} else {
				fmt.Println("Attempted to navigate to '", parts[2], "' from '", current.name, "' but was not present")
			}
		}
	}
}

func HandleNewDirectory(line string) {
	parts := strings.Split(line, " ")

	current.subDirectories = append(current.subDirectories, &Directory{parts[1], []int{}, []*Directory{}, current.name})
}

func HandleFile(line string) {
	parts := strings.Split(line, " ")

	size, _ := strconv.Atoi(parts[0])
	current.fileSizes = append(current.fileSizes, size)
}

func containsSubDirectory(subDirectoryName string) (bool, *Directory) {
	for _, subDirectory := range current.subDirectories {
		if subDirectory.name == subDirectoryName {
			return true, subDirectory
		}
	}

	return false, &Directory{}
}

func moveUpDirectory(parentName string) *Directory {
	return findDirectoryByName(parentName, top)
}

func findDirectoryByName(name string, position *Directory) *Directory {
	for _, directory := range position.subDirectories {
		if directory.name == name {
			return directory
		}

		findDirectoryByName(name, directory)
	}

	fmt.Println("No matching sub directory found for name '", name, "' in '", position.name, "'")
	return position
}

func sumDirectoriesUnderHundredThousand(directory *Directory) int {
	sum := 0
	size := getDirectorySize(directory)
	if size <= 100000 {
		sum += size
	}

	for _, sub := range directory.subDirectories {
		sum += sumDirectoriesUnderHundredThousand(sub)
	}

	return sum
}

func getDirectorySize(directory *Directory) int {
	size := 0
	for _, s := range directory.fileSizes {
		size += s
	}

	for _, sub := range directory.subDirectories {
		size += getDirectorySize(sub)
	}

	return size
}

func print(directory *Directory) {
	fmt.Println(directory.name, " ", directory.fileSizes)
	for _, subDirectory := range directory.subDirectories {
		print(subDirectory)
	}
}

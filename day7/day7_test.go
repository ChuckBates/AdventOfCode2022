package day7_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day7"
	"github.com/stretchr/testify/assert"
)

func TestComputePart1(t *testing.T) {
	if day7.ComputePart1("test.txt") != 95437 {
		t.Fatal("wrong value")
	}
}

func TestMoveUpOneDirectorySimple(t *testing.T) {
	top := &day7.Directory{"top", []int{}, []*day7.Directory{}, &day7.Directory{}}
	current := &day7.Directory{"current", []int{}, []*day7.Directory{}, top}
	top.SubDirectories = append(top.SubDirectories, current)

	day7.SetTopDirectory(top)
	day7.SetCurrentDirectory(current)

	result := day7.MoveUpDirectory()

	assert.Equal(t, top, result)
}

func TestMoveUpOneDirectoryComplex(t *testing.T) {
	directories := buildComplexStructure()
	day7.SetCurrentDirectory(directories[5])

	result := day7.MoveUpDirectory()

	assert.Equal(t, directories[4], result)
}

func buildComplexStructure() []*day7.Directory {
	top := &day7.Directory{"top", []int{}, []*day7.Directory{}, &day7.Directory{}}
	subDirectoryOne := &day7.Directory{"sub1", []int{}, []*day7.Directory{}, top}
	subDirectoryTwo := &day7.Directory{"sub2", []int{}, []*day7.Directory{}, top}
	top.SubDirectories = append(top.SubDirectories, subDirectoryOne)
	top.SubDirectories = append(top.SubDirectories, subDirectoryTwo)

	subDirectoryThree := &day7.Directory{"sub3", []int{}, []*day7.Directory{}, subDirectoryTwo}
	subDirectoryFour := &day7.Directory{"sub4", []int{}, []*day7.Directory{}, subDirectoryTwo}
	subDirectoryTwo.SubDirectories = append(subDirectoryTwo.SubDirectories, subDirectoryThree)
	subDirectoryTwo.SubDirectories = append(subDirectoryTwo.SubDirectories, subDirectoryFour)

	subDirectoryFive := &day7.Directory{"sub5", []int{}, []*day7.Directory{}, subDirectoryFour}
	subDirectoryFour.SubDirectories = append(subDirectoryFour.SubDirectories, subDirectoryFive)

	subDirectorySix := &day7.Directory{"sub6", []int{}, []*day7.Directory{}, subDirectoryOne}
	subDirectoryOne.SubDirectories = append(subDirectoryOne.SubDirectories, subDirectorySix)

	day7.SetTopDirectory(top)

	directories := []*day7.Directory{}
	directories = append(directories, top)
	directories = append(directories, subDirectoryOne)
	directories = append(directories, subDirectoryTwo)
	directories = append(directories, subDirectoryThree)
	directories = append(directories, subDirectoryFour)
	directories = append(directories, subDirectoryFive)
	directories = append(directories, subDirectorySix)

	return directories
}

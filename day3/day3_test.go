package day3_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day3"
)

func TestPart1(t *testing.T) {
	if day3.ComputePart1("test.txt") != 157 {
		t.Fatal("Wrong value")
	}
}
func TestPart2(t *testing.T) {
	if day3.ComputePart2("test.txt") != 70 {
		t.Fatal("Wrong value")
	}
}

func TestGetOrderedRucksacks(t *testing.T) {
	longestRucksack := "longest-rucksack"
	shortSackOne := "short-sack-1"
	shortSackTwo := "short-sack-2"
	expected := []string{longestRucksack, shortSackOne, shortSackTwo}

	result := day3.GetOrderedRucksacks([]string{shortSackOne, longestRucksack, shortSackTwo})
	if !arraysAreEqual(result, expected) {
		t.Fatal("Arrays are not equal")
	}
}

func arraysAreEqual(first []string, second []string) bool {
	return first[0] == second[0] && first[1] == second[1] && first[2] == second[2]
}

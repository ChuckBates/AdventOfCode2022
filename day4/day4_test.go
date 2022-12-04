package day4_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day4"
)

func TestPart1(t *testing.T) {
	if day4.ComputePart1("test.txt") != 2 {
		t.Fatal("Wrong value")
	}
}

func TestPart2(t *testing.T) {
	if day4.ComputePart2("test.txt") != 4 {
		t.Fatal("Wrong value")
	}
}

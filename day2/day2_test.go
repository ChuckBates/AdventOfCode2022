package day2_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day2"
)

func TestPart1(t *testing.T) {
	if day2.ComputePart1("test.txt") != 15 {
		t.Fatal("Wrong value")
	}
}

func TestPart2(t *testing.T) {
	if day2.ComputePart2("test.txt") != 12 {
		t.Fatal("Wrong value")
	}
}

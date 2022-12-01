package day1_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day1"
)

func TestPart1(t *testing.T) {
	if day1.ComputePart1() == 0 {
		t.Fatal("Wrong value")
	}
}

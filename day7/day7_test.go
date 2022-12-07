package day7_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day7"
)

func TestComputePart1(t *testing.T) {
	if day7.ComputePart1("test.txt") != 0 {
		t.Fatal("wrong value")
	}
}

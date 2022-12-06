package day6_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day6"
	"github.com/stretchr/testify/assert"
)

func TestComputePart1Input1(t *testing.T) {
	value := day6.ComputePart1("test1.txt")
	if value != 7 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart1Input2(t *testing.T) {
	value := day6.ComputePart1("test2.txt")
	if value != 5 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart1Input3(t *testing.T) {
	value := day6.ComputePart1("test3.txt")
	if value != 6 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart1Input4(t *testing.T) {
	value := day6.ComputePart1("test4.txt")
	if value != 10 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart1Input5(t *testing.T) {
	value := day6.ComputePart1("test5.txt")
	if value != 11 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart2Input1(t *testing.T) {
	value := day6.ComputePart2("test1.txt")
	if value != 19 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart2Input2(t *testing.T) {
	value := day6.ComputePart2("test2.txt")
	if value != 23 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart2Input3(t *testing.T) {
	value := day6.ComputePart2("test3.txt")
	if value != 23 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart2Input4(t *testing.T) {
	value := day6.ComputePart2("test4.txt")
	if value != 29 {
		t.Fatal("Wrong value", value)
	}
}

func TestComputePart2Input5(t *testing.T) {
	value := day6.ComputePart2("test5.txt")
	if value != 26 {
		t.Fatal("Wrong value", value)
	}
}

func TestQueueNotUnique(t *testing.T) {
	input := []string{"a", "b", "c", "b"}

	assert.False(t, day6.IsQueueUnique(input))
}

func TestQueueUnique(t *testing.T) {
	input := []string{"a", "b", "c", "d"}

	assert.True(t, day6.IsQueueUnique(input))
}

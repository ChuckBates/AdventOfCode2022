package day1_test

import (
	"testing"

	"github.com/ChuckBates/AdventOfCode2022/day1"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 71934, day1.ComputePart1("part1.txt"))
}

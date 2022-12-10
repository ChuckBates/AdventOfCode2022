package main

import (
	"fmt"

	"github.com/ChuckBates/AdventOfCode2022/day1"
	"github.com/ChuckBates/AdventOfCode2022/day10"
	"github.com/ChuckBates/AdventOfCode2022/day2"
	"github.com/ChuckBates/AdventOfCode2022/day3"
	"github.com/ChuckBates/AdventOfCode2022/day4"
	"github.com/ChuckBates/AdventOfCode2022/day5"
	"github.com/ChuckBates/AdventOfCode2022/day6"
	"github.com/ChuckBates/AdventOfCode2022/day7"
	"github.com/ChuckBates/AdventOfCode2022/day8"
	"github.com/ChuckBates/AdventOfCode2022/day9"
)

func main() {
	fmt.Println("Day 1 part 1: ", day1.ComputePart1("day1/part1.txt"))
	fmt.Println("Day 1 part 2: ", day1.ComputePart2("day1/part1.txt"))
	fmt.Println("Day 2 part 1: ", day2.ComputePart1("day2/input.txt"))
	fmt.Println("Day 2 part 2: ", day2.ComputePart2("day2/input.txt"))
	fmt.Println("Day 3 part 1: ", day3.ComputePart1("day3/input.txt"))
	fmt.Println("Day 3 part 2: ", day3.ComputePart2("day3/input.txt"))
	fmt.Println("Day 4 part 1: ", day4.ComputePart1("day4/input.txt"))
	fmt.Println("Day 4 part 2: ", day4.ComputePart2("day4/input.txt"))
	fmt.Println("Day 5 part 1: ", day5.ComputePart1("day5/input.txt"))
	fmt.Println("Day 5 part 2: ", day5.ComputePart2("day5/input.txt"))
	fmt.Println("Day 6 part 1: ", day6.ComputePart1("day6/input.txt"))
	fmt.Println("Day 6 part 2: ", day6.ComputePart2("day6/input.txt"))
	fmt.Println("Day 7 part 1: ", day7.ComputePart1("day7/input.txt"))
	fmt.Println("Day 7 part 2: ", day7.ComputePart2("day7/input.txt"))
	fmt.Println("Day 8 part 1: ", day8.ComputePart1("day8/input.txt"))
	fmt.Println("Day 8 part 2: ", day8.ComputePart2("day8/input.txt"))
	fmt.Println("Day 9 part 1: ", day9.ComputePart1("day9/input.txt"))
	fmt.Println("Day 9 part 2: ", day9.ComputePart2("day9/input.txt"))
	fmt.Println("Day 10 part 1: ", day10.ComputePart1("day10/input.txt"))
	day10.ComputePart2("day10/input.txt")
}

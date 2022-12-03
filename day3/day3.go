package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var priorityMap = make(map[string]int)

func ComputePart1(input string) int {
	buildPriorityMap()

	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	duplicates := []string{}
	newLineSeparator := regexp.MustCompile(`\n\s*`)
	values := newLineSeparator.Split(string(content), -1)
	for _, value := range values {
		duplicates = append(duplicates, getDuplicateItem(string(value)))
	}

	sum := 0
	for _, dupe := range duplicates {
		sum += getPriority(dupe)
	}
	return sum
}

func ComputePart2(input string) int {
	buildPriorityMap()

	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	badges := []string{}
	for _, group := range breakRucksacksIntoGroups(content) {
		badges = append(badges, getBadgeItem(group))
	}

	sum := 0
	for _, badge := range badges {
		sum += getPriority(badge)
	}
	return sum
}

func breakRucksacksIntoGroups(rucksacks []byte) [][]string {
	groups := [][]string{}

	newLineSeparator := regexp.MustCompile(`\n\s*`)
	values := newLineSeparator.Split(string(rucksacks), -1)
	sacksInGroup := []string{}
	for index, value := range values {
		if index != 0 && index%3 == 0 {
			groups = append(groups, sacksInGroup)
			sacksInGroup = []string{}
		}

		sacksInGroup = append(sacksInGroup, value)
	}
	groups = append(groups, sacksInGroup)

	return groups
}

func getBadgeItem(rucksackGroup []string) string {
	orderedRucksacks := GetOrderedRucksacks(rucksackGroup)
	rucksackOne := strings.Trim(orderedRucksacks[0], "\r")
	rucksackTwo := strings.Trim(orderedRucksacks[1], "\r")
	rucksackThree := strings.Trim(orderedRucksacks[2], "\r")

	for i := 0; i < len(rucksackOne); i++ {
		letter := rucksackOne[i]
		letterAsString := string(letter)
		if strings.Contains(rucksackTwo, letterAsString) && strings.Contains(rucksackThree, letterAsString) {
			return letterAsString
		}
	}

	fmt.Println("No badge item found!")
	return ""
}

func GetOrderedRucksacks(rucksackGroup []string) []string {
	longest, length := "", 0
	for _, rucksack := range rucksackGroup {
		if len(rucksack) > length {
			longest, length = rucksack, len(rucksack)
		}
	}

	result := []string{}
	result = append(result, longest)
	for _, rucksack := range rucksackGroup {
		if rucksack != longest {
			result = append(result, rucksack)
		}
	}

	return result
}

func getDuplicateItem(rucksack string) string {
	divisor := len(rucksack) / 2
	compartmentOne := rucksack[0:divisor]
	compartmentTwo := rucksack[divisor:]

	for i := 0; i < len(compartmentOne); i++ {
		if strings.Contains(compartmentTwo, string(compartmentOne[i])) {
			return string(compartmentOne[i])
		}
	}

	return ""
}

func getPriority(item string) int {
	value, containsKey := priorityMap[item]
	if !containsKey {
		fmt.Println("Map doesn't contain key ", item)
	}
	return value
}

func buildPriorityMap() {
	priorityMap["a"] = 1
	priorityMap["b"] = 2
	priorityMap["c"] = 3
	priorityMap["d"] = 4
	priorityMap["e"] = 5
	priorityMap["f"] = 6
	priorityMap["g"] = 7
	priorityMap["h"] = 8
	priorityMap["i"] = 9
	priorityMap["j"] = 10
	priorityMap["k"] = 11
	priorityMap["l"] = 12
	priorityMap["m"] = 13
	priorityMap["n"] = 14
	priorityMap["o"] = 15
	priorityMap["p"] = 16
	priorityMap["q"] = 17
	priorityMap["r"] = 18
	priorityMap["s"] = 19
	priorityMap["t"] = 20
	priorityMap["u"] = 21
	priorityMap["v"] = 22
	priorityMap["w"] = 23
	priorityMap["x"] = 24
	priorityMap["y"] = 25
	priorityMap["z"] = 26
	priorityMap["A"] = 27
	priorityMap["B"] = 28
	priorityMap["C"] = 29
	priorityMap["D"] = 30
	priorityMap["E"] = 31
	priorityMap["F"] = 32
	priorityMap["G"] = 33
	priorityMap["H"] = 34
	priorityMap["I"] = 35
	priorityMap["J"] = 36
	priorityMap["K"] = 37
	priorityMap["L"] = 38
	priorityMap["M"] = 39
	priorityMap["N"] = 40
	priorityMap["O"] = 41
	priorityMap["P"] = 42
	priorityMap["Q"] = 43
	priorityMap["R"] = 44
	priorityMap["S"] = 45
	priorityMap["T"] = 46
	priorityMap["U"] = 47
	priorityMap["V"] = 48
	priorityMap["W"] = 49
	priorityMap["X"] = 50
	priorityMap["Y"] = 51
	priorityMap["Z"] = 52
}

package day6

import (
	"strings"

	"github.com/ChuckBates/AdventOfCode2022/library"
)

func ComputePart1(input string) int {
	content := library.LoadFile(input)

	queue := make([]string, 0)
	for i := 0; i < len(content); i++ {
		queue = append(queue, string(content[i]))
		if len(queue) > 4 {
			queue = queue[1:]
			if IsQueueUnique(queue) {
				return i + 1
			}
		}
	}
	return 0
}

func ComputePart2(input string) int {
	content := library.LoadFile(input)

	queue := make([]string, 0)
	for i := 0; i < len(content); i++ {
		queue = append(queue, string(content[i]))
		if len(queue) > 14 {
			queue = queue[1:]
			if IsQueueUnique(queue) {
				return i + 1
			}
		}
	}
	return 0
}

func IsQueueUnique(queue []string) bool {
	queueConcat := strings.Join(queue, "")
	for _, value := range queue {
		if strings.Count(queueConcat, value) > 1 {
			return false
		}
	}
	return true
}

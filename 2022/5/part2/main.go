package main

import (
	"advent/helpers"
	"fmt"
	"strings"
)

func main() {
	lines, _ := helpers.ReadLines("../puzzle-input")

	moves := 0
	stacks := make(map[int][]string)

	for row, line := range lines {
		if line == "" {
			moves = row + 1
			break
		}

		for i, ch := range line {
			if ch > 'A' && ch <= 'Z' {
				offset := (i / 4) + 1
				stacks[offset] = append(stacks[offset], string(ch))
			}
		}
	}

	for i := 0; i < len(stacks); i++ {

		stacks[i] = reverse(stacks[i])
		s := make([]string, len(stacks[i]))
		copy(s, stacks[i])
	}

	for _, line := range lines[moves:] {

		count, from, to := parse_move(line)

		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-count:]...)
		stacks[from] = stacks[from][:len(stacks[from])-count]
	}

	for i := 1; i < len(stacks); i++ {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
}

func reverse(array []string) []string {
	for i := 0; i < len(array)/2; i++ {
		j := len(array) - i - 1
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func parse_move(move string) (int, int, int) {
	parts := strings.Split(move, " ")
	return helpers.Atoi(parts[1]), helpers.Atoi(parts[3]), helpers.Atoi(parts[5])
}

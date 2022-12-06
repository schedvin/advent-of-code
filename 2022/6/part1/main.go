package main

import (
	"advent/helpers"
	"fmt"
	"strings"
)

func duplicates(value string) bool {

	array := strings.Split(value, "")
	dup_map := make(map[string]int)

	for _, ch := range array {
		_, exist := dup_map[string(ch)]

		if !exist {
			dup_map[string(ch)] = 1
			continue
		}

		return true
	}

	return false
}

func main() {
	lines, _ := helpers.ReadLines("../puzzle-input")

	for _, line := range lines {
		for i := range line {

			if i-4 > 0 {
				begin := i - 4
				end := i

				slice := line[begin:end]
				fmt.Println(slice, " ", line)
				dup := duplicates(slice)

				if !dup {
					fmt.Println("Message Marker at ", i)
					break
				}
			}
		}
	}
}

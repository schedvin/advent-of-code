package main

import (
	"advent/helpers"
	"fmt"
	"strconv"
	"strings"
)

type SectionAssignment struct {
	start int
	end   int
}

func NewSectionAssignment(value []string) SectionAssignment {

	start, _ := strconv.Atoi(value[0])
	end, _ := strconv.Atoi(value[1])

	return SectionAssignment{
		start,
		end,
	}
}

func ParseSectionAssignments(value string) (SectionAssignment, SectionAssignment) {

	pair := strings.Split(value, ",")

	section1 := strings.Split(pair[0], "-")
	section2 := strings.Split(pair[1], "-")

	return NewSectionAssignment(section1), NewSectionAssignment(section2)
}

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

func main() {

	lines, _ := helpers.ReadLines("../puzzle-input")
	count := 0

	for _, line := range lines {
		first, second := ParseSectionAssignments(line)

		if InBetween(first.start, second.start, second.end) ||
			InBetween(first.end, second.start, second.end) {
			count++
			continue
		}

		if InBetween(second.start, first.start, first.end) ||
			InBetween(second.end, first.start, first.end) {
			count++
		}
	}

	fmt.Println(count)
}

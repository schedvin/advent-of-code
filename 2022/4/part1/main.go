package main

import (
	"advent/helpers"
	"fmt"
	"strconv"
	"strings"
)


type SectionAssignment struct {
  start int
  end int
}


func NewSectionAssignment(value []string)(SectionAssignment) {

  start, _ := strconv.Atoi(value[0])
  end, _ := strconv.Atoi(value[1])

  return SectionAssignment { 
    start, 
    end,
  }
}

func ParseSectionAssignments(value string)(SectionAssignment, SectionAssignment) {

  pair := strings.Split(value, ",")

  section1 := strings.Split(pair[0], "-")
  section2 := strings.Split(pair[1], "-")

  return NewSectionAssignment(section1), NewSectionAssignment(section2)
}

func main() {

  lines, _ := helpers.ReadLines("../puzzle-input")
  count := 0

  for _, line := range lines {

    first, second := ParseSectionAssignments(line)

    if first.start <= second.start && first.end >= second.end {
      count++
      continue
    }

    if (second.start <= first.start && second.end >= first.end) {
      count++
    }
  }

  fmt.Println(count)
}

package main

import (
	"fmt"
	"strconv"
  "advent/helpers"
)

func main() {

  lines, _ := helpers.ReadLines("./puzzle-input")

  var max int

  total := 0

  for _, line := range lines {
    if line == "" {
      total = 0;
      continue;
    }
    
    temp,_  := strconv.Atoi(line)
    total += temp
    
    if total >= max {
      max = total
    } 
  }

  result := fmt.Sprintf("Total calories carried by one Elf: %d calories", max)
  fmt.Println(result)
}

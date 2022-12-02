
package main

import (
	"fmt"
  "sort"
	"strconv"
  "advent/helpers"
)


func main() {

  lines, _ := helpers.ReadLines("../puzzle-input")

  total := 0
  var totals []int

  for _, line := range lines {
    if line == "" {
      totals = append(totals, total)
      total = 0;
      continue;
    }
    
    temp,_  := strconv.Atoi(line)
    total += temp
  }

  sort.Sort(sort.Reverse(sort.IntSlice(totals)))

  top := totals[0:3]
  sum := helpers.Sum(top)

  result := fmt.Sprintf("Total calories carried by top 3 Elves: %d calories", sum)
  fmt.Println(result)
}


package main

import (
	"fmt"
  "strings"
	
  "advent/helpers"
)

func NewSlice(start, end, step int) []int {
    if step <= 0 || end < start {
        return []int{}
    }
    s := make([]int, 0, 1+(end-start)/step)
    for start <= end {
        s = append(s, start)
        start += step
    }
    return s
}

func RemoveDuplicate(array []string) []string {
    m := make(map[string]string)
    for _, x := range array {
        m[x] = x
    }
    var ClearedArr []string
    for x, _ := range m {
        ClearedArr = append(ClearedArr, x)
    }
    return ClearedArr
}

func main() {

  priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

  lines, _ := helpers.ReadLines("../puzzle-input")
  subTotal := 0
  total := 0

  indexes := NewSlice(2, len(lines) -1, 3)
  
  for _, val := range indexes {

    row1 := strings.Split(lines[val - 1], "")
    row2 := strings.Split(lines[val - 2], "")
    row3 := strings.Split(lines[val], "")

    row1 = RemoveDuplicate(row1)
    row2 = RemoveDuplicate(row2)
    row3 = RemoveDuplicate(row3)

    group1 := strings.Join(row1, "")
    group2 := strings.Join(row2, "")
    group3 := strings.Join(row3, "")

    for _, char := range group3 {
      if strings.IndexAny(group1, string(char)) > -1 &&
        strings.IndexAny(group2, string(char)) > -1 {
        subTotal += strings.Index(priorities, string(char)) + 1       
      }
    }
    total += subTotal
    subTotal = 0
  } 

  fmt.Println("Sum of priorities for for items types of all groups: ", total)
}

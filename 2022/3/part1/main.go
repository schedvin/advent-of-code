package main

import (
	"advent/helpers"
	"fmt"
  "strings"
)

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
  total := 0

  for _, line := range lines {
    
    length := len([]rune(line))
    half := length / 2

    first := strings.Join(RemoveDuplicate(strings.Split(line[0 : half], "")), "")    
    last := strings.Join(RemoveDuplicate(strings.Split(line[half : length], "")), "")

    for _, char := range first {
      if strings.IndexAny(last, string(char)) > -1 {
        total += strings.Index(priorities, string(char)) + 1
      } 
    }
  }

  fmt.Println(total)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func shouldPanic(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  file, err := os.Open("./puzzle-input")

  shouldPanic(err)

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var lines []string
  var max int

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

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

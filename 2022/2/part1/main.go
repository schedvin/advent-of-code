package main

import (
  "fmt"
  "strings"

  "advent/helpers"
)

func Splair(val string)(string, string) {

  pair := strings.Split(val, " ");

  return pair[0], pair[1]
}

func main() {
 
  const LOSS = 0
  const WIN = 6
  const DRAW = 3

  lines, _ := helpers.ReadLines("../puzzle-input")
  score := 0
  
  points := make(map[string]int)

  // rock
  points["X"] = 1
  // paper
  points["Y"] = 2
  // scissors
  points["Z"] = 3

  // rock
  points["A"] = 1
  // paper
  points["B"] = 2
  // scissors
  points["C"] = 3

  x := "X"
  y := "Y"
  z := "Z"

  fmt.Println("X", points[x])
  fmt.Println("Y", points[y])
  fmt.Println("Z", points[z])

  for _, val := range lines {
    opponent, self := Splair(val)

    // draw
    if points[opponent] == points[self] {
      score += points[self] + DRAW
      continue
    }
   
    switch self {
      case "X":
        if opponent == "B" {
          score += points[self] + LOSS
          continue
        }
      case "Y":
        if opponent == "C" {
          score += points[self] + LOSS
          continue
        }
      case "Z":
        if (opponent == "A") {
          score += points[self] + LOSS
          continue
        }
    }

    score += points[self] + WIN
  }

  fmt.Println("Score : ", score)
}

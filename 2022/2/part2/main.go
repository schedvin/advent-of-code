
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

type play struct {
  win string
  loss string
}

func main() {

  playbook := make(map[string]play)

  playbook["A"] = play { "Y", "Z" }
  playbook["B"] = play { "Z", "X" }
  playbook["C"] = play { "X", "Y" }

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

  for _, val := range lines {
    opponent, play := Splair(val)

    strategy := playbook[opponent]
    self := ""

    switch play {
      case "X":
        self = strategy.loss
      case "Y":
        self = opponent
      case "Z":
        self = strategy.win
    }

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

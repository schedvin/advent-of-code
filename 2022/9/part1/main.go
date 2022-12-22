package main

import (
	"advent/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func MarkVisited(visited map[Position]bool, p Position) {
	if _, exist := visited[p]; !exist {
		visited[p] = true
	}
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func main() {
	lines, _ := helpers.ReadLines("./puzzle-input")

	head := Position{x: 0, y: 0}
	tail := Position{x: 0, y: 0}
	visited := make(map[Position]bool)

	for _, line := range lines {

		move := strings.Split(line, " ")

		direction := move[0]
		distance, _ := strconv.Atoi(move[1])

		for i := 0; i < distance; i++ {
			if direction == "U" {
				head.y += 1

				if Abs(head.y-tail.y) > 1 &&
					Abs(head.x-tail.x) >= 1 {
					tail.y += 1
					tail.x = head.x
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
					continue
				}

				if Abs(head.y-tail.y) > 1 {
					tail.y += 1
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
				}
			}
			if direction == "D" {
				head.y -= 1

				if Abs(head.y-tail.y) > 1 &&
					Abs(head.x-tail.x) >= 1 {
					tail.y -= 1
					tail.x = head.x
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
					continue
				}
				if Abs(head.y-tail.y) > 1 {
					tail.y -= 1
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
				}

			}
			if direction == "L" {
				head.x -= 1

				if Abs(head.x-tail.x) > 1 &&
					Abs(head.y-tail.y) >= 1 {
					tail.y = head.y
					tail.x -= 1
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
					continue
				}

				if Abs(head.x-tail.x) > 1 {
					tail.x -= 1
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
				}

			}
			if direction == "R" {
				head.x += 1

				if Abs(head.x-tail.x) > 1 &&
					Abs(head.y-tail.y) >= 1 {
					tail.y = head.y
					tail.x += 1
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
					continue
				}

				if Abs(head.x-tail.x) > 1 {
					tail.x += 1
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
				}
			}

			fmt.Println(head, tail)
		}
	}
	fmt.Println(len(visited))
}

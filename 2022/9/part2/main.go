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

func FollowKnot(direction string, next Position, p *Position) {

	if direction == "U" {

		if Abs(next.y-p.y) > 1 &&
			Abs(next.x-p.x) >= 1 {
			p.y += 1
			p.x = next.x
			return
		}

		if Abs(next.y-p.y) > 1 {
			p.y += 1
		}
	}
	if direction == "D" {

		if Abs(next.y-p.y) > 1 &&
			Abs(next.x-p.x) >= 1 {
			p.y -= 1
			p.x = next.x
			return
		}
		if Abs(next.y-p.y) > 1 {
			p.y -= 1
		}

	}
	if direction == "L" {

		if Abs(next.x-p.x) > 1 &&
			Abs(next.y-p.y) >= 1 {
			p.y = next.y
			p.x -= 1
			return
		}

		if Abs(next.x-p.x) > 1 {
			p.x -= 1
		}

	}
	if direction == "R" {
		if Abs(next.x-p.x) > 1 &&
			Abs(next.y-p.y) >= 1 {
			p.y = next.y
			p.x += 1
			return
		}

		if Abs(next.x-p.x) > 1 {
			p.x += 1
		}
	}
}

func main() {
	lines, _ := helpers.ReadLines("../puzzle-input")

	head := Position{x: 0, y: 0}
	tail := Position{x: 0, y: 0}
	rope := []Position{
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
		Position{x: 0, y: 0},
	}
	visited := make(map[Position]bool)

	for _, line := range lines {

		move := strings.Split(line, " ")

		direction := move[0]
		distance, _ := strconv.Atoi(move[1])

		for i := 0; i < distance; i++ {
			temp_tail := Position{tail.x, tail.y}
			if direction == "U" {
				head.y += 1
				FollowKnot(direction, head, &rope[0])
				for j := 1; j < len(rope); j++ {
					FollowKnot(direction, rope[j-1], &rope[j])
				}
				FollowKnot(direction, rope[len(rope)-1], &tail)

				if temp_tail.x != tail.x || temp_tail.y != temp_tail.y {
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
				}
			}
			if direction == "D" {
				head.y -= 1

				FollowKnot(direction, head, &rope[0])
				for j := 1; j < len(rope); j++ {
					FollowKnot(direction, rope[j-1], &rope[j])
				}
				FollowKnot(direction, rope[len(rope)-1], &tail)

				if temp_tail.x != tail.x || temp_tail.y != temp_tail.y {
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
				}
			}
			if direction == "L" {
				head.x -= 1

				FollowKnot(direction, head, &rope[0])
				for j := 1; j < len(rope); j++ {
					FollowKnot(direction, rope[j-1], &rope[j])
				}

				FollowKnot(direction, rope[len(rope)-1], &tail)

				if temp_tail.x != tail.x || temp_tail.y != temp_tail.y {
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
				}
			}
			if direction == "R" {
				head.x += 1

				FollowKnot(direction, head, &rope[0])
				for j := 1; j < len(rope); j++ {
					FollowKnot(direction, rope[j-1], &rope[j])
				}
				FollowKnot(direction, rope[len(rope)-1], &tail)

				if temp_tail.x != tail.x || temp_tail.y != temp_tail.y {
					MarkVisited(visited, Position{x: tail.x, y: tail.y})
				}
			}
			fmt.Println(head, rope, tail)
		}
	}
	fmt.Println(len(visited))
}

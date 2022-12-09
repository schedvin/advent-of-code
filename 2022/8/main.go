package main

import (
	"advent/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	lines, _ := helpers.ReadLines("./puzzle-input")
	forest := [][]int{}

	for row, line := range lines {
		forest = append(forest, []int{})
		trees := strings.Split(line, "")

		for _, tree := range trees {
			height, _ := strconv.Atoi(tree)
			forest[row] = append(forest[row], height)
		}
	}

	// length
	length := len(forest)
	width := len(forest[0])

	sum := length * 2
	// width
	sum += width * 2
	// corners
	sum -= 4

	for i := 1; i < length-1; i++ {
		for j := 1; j < width-1; j++ {

			x1 := forest[i][0:j]
			x2 := forest[i][j+1:]

			y1 := forest[0:i]
			y2 := forest[i+1:]

			fmt.Println(getColumn(y2, j))
			if greaterThan(x1, forest[i][j]) || greaterThan(getColumn(y1, j), forest[i][j]) {
				sum += 1
				continue
			}

			if greaterThan(x2, forest[i][j]) || greaterThan(getColumn(y2, j), forest[i][j]) {
				sum += 1
				continue
			}
		}
	}

	fmt.Println(sum)
}

func getColumn(values [][]int, column int) []int {
	result := []int{}
	for _, v := range values {
		result = append(result, v[column])
	}
	return result
}

func greaterThan(values []int, value int) bool {
	for _, e := range values {
		if value <= e {
			return false
		}
	}
	return true
}

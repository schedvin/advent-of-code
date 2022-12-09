package main

import (
	"advent/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	lines, _ := helpers.ReadLines("../puzzle-input")
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

	score := 0

	for i := 1; i < length-1; i++ {
		for j := 1; j < width-1; j++ {

			x1 := reverse(forest[i][:j])
			x2 := forest[i][j+1:]

			y1 := forest[:i]
			y2 := forest[i+1:]

			treeHeight := forest[i][j]

			left := getDistance(x1, treeHeight)
			right := getDistance(x2, treeHeight)

			fmt.Println(y1)
			top := getDistance(getColumn(y1, j), treeHeight)
			bottom := getDistance(getBottomColumn(y2, j), treeHeight)

			sum := left * right * top * bottom

			//fmt.Println(sum, " = ", left, " * ", right, " * ", top, " * ", bottom)
			if sum > score {
				score = sum
			}

		}
	}

	fmt.Println("Score: ", score)

}

func reverse(array []int) []int {
	result := []int{}

	for _, v := range array {
		result = insert(result, v, 0)
	}

	return result
}
func insert(array []int, element int, i int) []int {
	return append(array[:i], append([]int{element}, array[i:]...)...)
}

func getColumn(values [][]int, column int) []int {
	result := []int{}
	for _, v := range values {
		result = insert(result, v[column], 0)
	}
	return result
}

func getBottomColumn(values [][]int, column int) []int {
	result := []int{}
	for _, v := range values {
		result = append(result, v[column])
	}
	return result
}

func getDistance(trees []int, treeHeight int) int {
	for k, height := range trees {
		if treeHeight <= height {
			return k + 1
		}
	}
	return len(trees)
}

package main

import (
	"advent/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, _ := helpers.ReadLines("../puzzle-input")

	paths := []string{}
	files := make(map[string]int)
	cwd := []string{}

	for _, line := range lines {

		tokens := strings.Split(line, " ")

		// command
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					cwd = cwd[:len(cwd)-1]
					continue
				}

				if tokens[2] != "/" {
					cwd = append(cwd, tokens[2])
					paths = append(paths, strings.Join(cwd, "/"))
				}
			}
			continue
		}

		if tokens[0] == "dir" {
			continue
		}

		size, _ := strconv.Atoi(tokens[0])
		files[strings.Join(cwd, "/")+"/"+tokens[1]] = size
	}

	usedSpace := 0
	for _, size := range files {
		usedSpace += size
	}

	capacity := 70000000
	freeSpace := capacity - usedSpace
	requiredSpace := 30000000 - freeSpace
	minimumSpace := int(^uint(0) >> 1)

	solution1 := 0

	for _, path := range paths {
		size := 0
		for file, fileSize := range files {
			if strings.HasPrefix(file, path) {
				size += fileSize
			}
		}
		if size <= 100000 {
			solution1 += size
		}

		if size >= requiredSpace && size < minimumSpace {
			minimumSpace = size
		}
	}

	fmt.Println("Total size: ", solution1)
	fmt.Println("Minimum space: ", minimumSpace)
}

package main

import (
	"fmt"
	"strings"
)

func lights2vector(lights string) []int {
	vector := make([]int, len(lights)-2)
	for i := 1; i < len(lights)-1; i++ {
		if lights[i] == 46 { // "." = 46
			vector[i-1] = 0
		}
		if lights[i] == 35 { // "#" = 35
			vector[i-1] = 1
		}
	}
	return vector
}

func day10part1(filename string) {
	input := FileToArray(filename)

	for _, row := range input {
		parts := strings.Split(row, " ")
		// parts[0] 			lights
		// parts[...]			vectors (raw)
		// parts[len(parts)-1]	joltage

		vectorLights := lights2vector(parts[0])
		fmt.Println(vectorLights)

		for i := 1; i < len(vectorLights)-1; i++ {

		}
	}
	fmt.Println("day10")
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func File2Matrix(filename string) [496][2]int {
	input := FileToArray(filename)

	var matrix [496][2]int

	for i, row := range input {
		rowSplit := strings.Split(row, ",")
		for j, col := range rowSplit {
			matrix[i][j], _ = strconv.Atoi(col)
		}
	}
	return matrix
}

func day9part1(filename string) {
	matrix := File2Matrix(filename)

	for _, row := range matrix {
		fmt.Println(row)
	}

	areaMax := 0
	for i := 0; i < len(matrix); i++ { // base = matrix[i]
		for j := i + 1; j < len(matrix); j++ {
			// base := matrix[i]
			// corner := matrix[j]
			// fmt.Println(base, corner)
			width := Abs(matrix[i][0]-matrix[j][0]) + 1
			height := Abs(matrix[i][1]-matrix[j][1]) + 1
			area := width * height
			if area > areaMax {
				areaMax = area
			}
		}
	}

	fmt.Println(areaMax)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// [8][2]int
// [496][2]int

func File2Matrix(filename string) [][2]int {
	input := FileToArray(filename)

	coordinates := make([][2]int, len(input))

	for i, row := range input {
		rowSplit := strings.Split(row, ",")
		for j, col := range rowSplit {
			coordinates[i][j], _ = strconv.Atoi(col)
		}
	}
	return coordinates
}

// dev matrix: [9][14]string
// prd matrix: [98400][98300]string
// or from highest coordinates: [98328][98229]string

func addInBetweenCoordinates(edgeCoordinates [][2]int, inBetweenCoordinates *[][2]int) {

	for i := 0; i < len(edgeCoordinates); i++ { // base = coordinates[i]
		for j := i + 1; j < len(edgeCoordinates); j++ {
			// Horizontal
			if edgeCoordinates[i][0] == edgeCoordinates[j][0] {
				// add green coordinates / draw "G"
				lowest, highest := Order(edgeCoordinates[i][1], edgeCoordinates[j][1])
				for k := lowest + 1; k < highest; k++ {
					*inBetweenCoordinates = append(*inBetweenCoordinates, [2]int{edgeCoordinates[j][0], k})
				}
			}
			// vertical
			if edgeCoordinates[i][1] == edgeCoordinates[j][1] {
				lowest, highest := Order(edgeCoordinates[i][0], edgeCoordinates[j][0])
				for k := lowest + 1; k < highest; k++ {
					*inBetweenCoordinates = append(*inBetweenCoordinates, [2]int{k, edgeCoordinates[j][1]})
				}
			}
		}
	}
}

//go:noinline
func slice2map(coordinates [][2]int) map[[2]int]bool {
	coordinateMap := make(map[[2]int]bool, len(coordinates))
	for i := 0; i < len(coordinates); i++ {
		row := coordinates[i]
		coordinateMap[row] = true
	}
	return coordinateMap
}

func drawCoordinates(redMap map[[2]int]bool, greenMap map[[2]int]bool, row int, column int) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			character2print := "."
			if greenMap[[2]int{j, i}] {
				character2print = "G"
			}
			if redMap[[2]int{j, i}] { // may overwrite a "G"
				character2print = "R"
			}
			fmt.Print(character2print)
			if j == column-1 {
				fmt.Println()
			}
		}
	}
}

func day9part2(filename string) {
	redCoordinates := File2Matrix(filename)
	//matrix (9, 14) // dev
	//matrix (98400, 98300) // prd

	// set Green boundaries
	greenCoordinates := make([][2]int, 0)

	addInBetweenCoordinates(redCoordinates, &greenCoordinates)
	addInBetweenCoordinates(greenCoordinates, &greenCoordinates)
	fmt.Println(len(greenCoordinates))

	// convert to maps for faster "contains"
	redMap := slice2map(redCoordinates)
	fmt.Println(redMap)
	greenMap := slice2map(greenCoordinates)
	fmt.Println(greenMap)
	fmt.Println(len(greenMap))

	drawCoordinates(redMap, greenMap, 9, 14)

	areaMax := 0
	for i := 0; i < len(redCoordinates); i++ { // base = coordinates[i]
		for j := i + 1; j < len(redCoordinates); j++ {
			// base := coordinates[i]
			// corner := coordinates[j]
			// fmt.Println(base, corner)
			width := Abs(redCoordinates[i][0]-redCoordinates[j][0]) + 1
			height := Abs(redCoordinates[i][1]-redCoordinates[j][1]) + 1
			area := width * height
			if area > areaMax {
				// if all corners are in red or green it is ok, is it?
				oppositeCorner1 := [2]int{redCoordinates[i][1], redCoordinates[j][0]}
				oppositeCorner2 := [2]int{redCoordinates[i][0], redCoordinates[j][1]}
				if (greenMap[oppositeCorner1] || redMap[oppositeCorner1]) &&
					(!greenMap[oppositeCorner2] || !redMap[oppositeCorner2]) {
					areaMax = area
					fmt.Println(areaMax)
				}
			}
		}
	}

	fmt.Println("92981 < ", areaMax, " < 4499879904")

}

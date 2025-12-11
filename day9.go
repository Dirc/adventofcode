package main

import (
	"fmt"
	"strconv"
	"strings"
)

// [8][2]int
// [496][2]int

func File2Matrix(filename string) []Coordinate {
	input := FileToArray(filename)

	coordinates := make([]Coordinate, len(input))

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

func addInBetweenCoordinates(edgeCoordinates []Coordinate, inBetweenCoordinates *[]Coordinate) {

	for i := 0; i < len(edgeCoordinates); i++ { // base = coordinates[i]
		for j := i + 1; j < len(edgeCoordinates); j++ {
			// Horizontal
			if edgeCoordinates[i][0] == edgeCoordinates[j][0] {
				// add green coordinates / draw "G"
				lowest, highest := Order(edgeCoordinates[i][1], edgeCoordinates[j][1])
				for k := lowest + 1; k < highest; k++ {
					*inBetweenCoordinates = append(*inBetweenCoordinates, Coordinate{edgeCoordinates[j][0], k})
				}
			}
			// vertical
			if edgeCoordinates[i][1] == edgeCoordinates[j][1] {
				lowest, highest := Order(edgeCoordinates[i][0], edgeCoordinates[j][0])
				for k := lowest + 1; k < highest; k++ {
					*inBetweenCoordinates = append(*inBetweenCoordinates, Coordinate{k, edgeCoordinates[j][1]})
				}
			}
		}
	}
}

type Coordinate [2]int

//go:noinline
func slice2map(coordinates []Coordinate) map[Coordinate]bool {
	lenght := len(coordinates)
	coordinateMap := make(map[Coordinate]bool, lenght)
	// var coordinateMap map[Coordinate]bool
	for i := 0; i < lenght; i++ {
		row := coordinates[i]
		coordinateMap[row] = true
	}
	return coordinateMap
}

func drawCoordinates(redMap map[Coordinate]bool, greenMap map[Coordinate]bool, row int, column int) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			character2print := "."
			if greenMap[Coordinate{j, i}] {
				character2print = "G"
			}
			if redMap[Coordinate{j, i}] { // may overwrite a "G"
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
	// matrix (9, 14) // dev
	// matrix (98400, 98300) // prd

	// set Green boundaries
	greenCoordinates := make([]Coordinate, 0)

	addInBetweenCoordinates(redCoordinates, &greenCoordinates)
	fmt.Println(len(greenCoordinates))
	addInBetweenCoordinates(greenCoordinates, &greenCoordinates)
	fmt.Println(len(greenCoordinates))

	// convert to maps for faster "contains"
	redMap := slice2map(redCoordinates)
	fmt.Println(redMap)
	greenMap := slice2map(greenCoordinates)
	fmt.Println(greenMap)
	fmt.Println(len(greenMap))

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
				// Check: corners only
				// if all corners are in red or green it is ok, is it?
				// oppositeCorner1 := Coordinate{redCoordinates[i][1], redCoordinates[j][0]}
				// oppositeCorner2 := Coordinate{redCoordinates[i][0], redCoordinates[j][1]}
				// if (greenMap[oppositeCorner1] || redMap[oppositeCorner1]) &&
				// 	(greenMap[oppositeCorner2] || redMap[oppositeCorner2]) {
				// 	areaMax = area
				// 	fmt.Println(areaMax)
				// }

				// Check: full interior
				smallest_r, biggest_r := Order(redCoordinates[i][0], redCoordinates[j][0])
				smallest_c, biggest_c := Order(redCoordinates[i][1], redCoordinates[j][1])
			rgCheckLoop:
				for m := smallest_r; m < biggest_r; m++ {
					for n := smallest_c; n < biggest_c; n++ {
						if !greenMap[Coordinate{m, n}] && !greenMap[Coordinate{m, n}] {
							fmt.Println("Interior coordinate not red/green: ", m, n)
							fmt.Println("For area with corners: ", redCoordinates[i], redCoordinates[j])
							break rgCheckLoop
						}
						areaMax = area
						fmt.Println(areaMax, "with coordinates: ", redCoordinates[i], redCoordinates[j])
					}
				}
			}
		}
	}

	drawCoordinates(redMap, greenMap, 9, 14)
	fmt.Println("92981 < ", areaMax, " < 4499879904")

}

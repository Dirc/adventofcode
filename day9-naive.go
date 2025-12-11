package main

//
// import (
//	"fmt"
//	"strconv"
//	"strings"
// )
//
// // [8][2]int
// // [496][2]int
//
// func File2Matrix(filename string) [496][2]int {
//	input := FileToArray(filename)
//
//	coordinates := [496][2]int{}
//
//	for i, row := range input {
//		rowSplit := strings.Split(row, ",")
//		for j, col := range rowSplit {
//			coordinates[i][j], _ = strconv.Atoi(col)
//		}
//	}
//	return coordinates
// }
//
// func day9part1(filename string) {
//	coordinates := File2Matrix(filename)
//
//	for _, row := range coordinates {
//		fmt.Println(row)
//	}
//
//	areaMax := 0
//	for i := 0; i < len(coordinates); i++ { // base = coordinates[i]
//		for j := i + 1; j < len(coordinates); j++ {
//			// base := coordinates[i]
//			// corner := coordinates[j]
//			// fmt.Println(base, corner)
//			width := Abs(coordinates[i][0]-coordinates[j][0]) + 1
//			height := Abs(coordinates[i][1]-coordinates[j][1]) + 1
//			area := width * height
//			if area > areaMax {
//				areaMax = area
//			}
//		}
//	}
//
//	fmt.Println(areaMax)
// }
//
// // dev matrix: [9][14]string
// // prd matrix: [98400][98300]string
// // or from highest coordinates: [98328][98229]string
//
// func drawBetweenHorizontalPoints(matrix *[][]string, between string, draw string) {
//	// set Green boundaries
//	// horizontal boundaries
//	for i, row := range *matrix {
//		countRed := 0
//		var greenLine [2]int
//		for j := 0; j < len(row); j++ {
//			if row[j] == between && countRed == 0 {
//				countRed++
//				greenLine[0] = j
//			}
//			if row[j] == between && countRed >= 1 {
//				countRed++
//				greenLine[1] = j
//			}
//		}
//		if countRed >= 2 {
//			// draw "G"
//			for k := greenLine[0] + 1; k < greenLine[1]; k++ {
//				matrix[i][k] = draw
//			}
//		}
//	}
// }
//
// func drawBetweenVerticalPoints(matrix *[][]string, between string, draw string) {
//	// vertical boundaries
//	for j := 0; j < len(matrix[0]); j++ {
//		countRed := 0
//		var greenLine [2]int
//		for i := 0; i < len(matrix); i++ {
//			if matrix[i][j] == between && countRed == 0 {
//				countRed++
//				greenLine[0] = i
//			}
//			if matrix[i][j] == between && countRed >= 1 {
//				countRed++
//				greenLine[1] = i
//			}
//		}
//		if countRed >= 2 {
//			// draw "G"
//			for k := greenLine[0] + 1; k < greenLine[1]; k++ {
//				matrix[k][j] = draw
//			}
//		}
//	}
// }
//
// func createMatrix(coordinates [496][2]int, rows int, columns int) [][]string {
//	matrix := make([][]string, rows)
//	for i := range matrix {
//		matrix[i] = make([]string, columns)
//		for j := range matrix[i] {
//			matrix[i][j] = "."
//		}
//	}
//	//drawMatrix(matrix)
//
//	// setRed
//	for _, row := range coordinates {
//		matrix[row[1]][row[0]] = "R"
//	}
//	//drawMatrix(matrix)
//
//	drawBetweenHorizontalPoints(&matrix, "R", "G")
//	//drawMatrix(matrix)
//
//	drawBetweenVerticalPoints(&matrix, "R", "G")
//	//drawMatrix(matrix)
//
//	// Set Green filling
//	// for each point, check if left/right, up/down all have G -> then write G
//	// horizontal
//	drawBetweenHorizontalPoints(&matrix, "G", "G")
//	//drawMatrix(matrix)
//
//	drawBetweenVerticalPoints(&matrix, "G", "G")
//	//drawMatrix(matrix)
//
//	// setRed again
//	for _, row := range coordinates {
//		matrix[row[1]][row[0]] = "R"
//	}
//	//drawMatrix(matrix)
//
//	return matrix
// }
//
// func drawMatrix(matrix [][]string) {
//	fmt.Println("")
//	for i := 0; i < len(matrix); i++ {
//		fmt.Println(matrix[i])
//	}
//	fmt.Println("")
// }
//
// func GetHighest(filename string) {
//	high0 := 0
//	high1 := 0
//	coordinates := File2Matrix(filename)
//
//	for _, row := range coordinates {
//		if row[0] > high0 {
//			high0 = row[0]
//		}
//		if row[1] > high1 {
//			high1 = row[1]
//		}
//	}
//	fmt.Println(high0, high1)
// }
//
// func day9part2(filename string) {
//	coordinates := File2Matrix(filename)
//	//matrix := createMatrix(coordinates, 9, 14) // dev
//	matrix := createMatrix(coordinates, 98400, 98300) // prd
//
//	areaMax := 0
//	//outerloop:
//	for i := 0; i < len(coordinates); i++ { // base = coordinates[i]
//		//secondloop:
//		for j := i + 1; j < len(coordinates); j++ {
//			// base := coordinates[i]
//			// corner := coordinates[j]
//			// fmt.Println(base, corner)
//			width := Abs(coordinates[i][0]-coordinates[j][0]) + 1
//			height := Abs(coordinates[i][1]-coordinates[j][1]) + 1
//			area := width * height
//			if area > areaMax {
//				// check interior for dottes
// 				dotsFound := false
// 				smallest_r, biggest_r := Order(coordinates[i][1], coordinates[j][1])
// 				smallest_c, biggest_c := Order(coordinates[i][0], coordinates[j][0])
// 			dotCheckLoop:
// 				for m := smallest_r + 1; m < biggest_r; m++ {
// 					for n := smallest_c + 1; n < biggest_c; n++ {
// 						if matrix[m][n] == "." {
// 							dotsFound = true
// 							break dotCheckLoop
// 						}
// 					}
// 				}
//				if !dotsFound {
//					areaMax = area
//				}
//			}
//		}
//	}
//
//	fmt.Println(areaMax)
//
// }

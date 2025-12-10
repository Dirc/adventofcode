package main

import "fmt"

func day7part1(filename string) {
	matrix := FileToArray(filename)

	// 46 = .
	// 83 = S (start beam)
	// 94 = ^ (beam splitter)

	directions := [][2]int{
		[2]int{-1, 0},  // move
		[2]int{-1, -1}, // split left
		[2]int{-1, 1},  // split right
	}

	// design:
	// remove dotted rows
	// for each line:
	// - determine positions next line
	// - check them and make a new list
	//   - make it unique + sort
	// replay
	var matrixUndotted []string
	for _, row := range matrix {
		if row != "..............." {
			matrixUndotted = append(matrixUndotted, row)
		}
	}
	DrawStringArray(matrixUndotted)

	fmt.Println(directions[0])

	fmt.Println(matrix[1][1])
	fmt.Println(matrix[0][7])
	fmt.Println(matrix[2][7])

	fmt.Println(matrix[4][6])

}

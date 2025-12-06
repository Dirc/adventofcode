package main

import (
	"fmt"
)

func checkPaperRole(apenstaartje rune, runes []rune, i int) bool {
	if i < 0 || i >= len(runes) {
		return false
	}
	if runes[i] == apenstaartje {
		return true
	}
	return false
}

func draw(xes []int, original []rune) {
	printString := ""
	for i := 0; i < len(original); i++ {
		if i%137 == 0 {
			printString += "\n"
		}
		if Contains(xes, i) {
			printString += "x"
		} else {
			printString += string(original[i])
		}
	}
	fmt.Println(printString)
}

func day4Part1(filename string) {
	var freePaperRoles []int
	var totalFree int
	//var width int

	input := fileToArray(filename)

	var oneline string
	for _, line := range input {
		oneline += line
	}
	fmt.Println(oneline)
	onelineRunes := []rune(oneline)

	apenstaartje := onelineRunes[2]

	for i := 0; i < len(onelineRunes); i++ {
		if onelineRunes[i] != apenstaartje {
			continue
		}

		var countAdjecentRoles int

		indexList := [8]int{i - 137 - 1, i - 137, i - 137 + 1, i - 1, i + 1, i + 137 - 1, i + 137, i + 137 + 1}
		leftBorder := [5]int{i - 137, i - 137 + 1, i + 1, i + 137, i + 137 + 1}
		rightBorder := [5]int{i - 137 - 1, i - 137, i - 1, i + 137 - 1, i + 137}

		// if on left border
		if i%137 == 0 {
			for _, j := range leftBorder {
				if checkPaperRole(apenstaartje, onelineRunes, j) {
					countAdjecentRoles++
				}
			}
		}

		// if on right border
		if i%137 == 137-1 {
			for _, j := range rightBorder {
				if checkPaperRole(apenstaartje, onelineRunes, j) {
					countAdjecentRoles++
				}
			}
		}

		// Not on border
		if i%137 != 0 && i%137 != 137-1 {
			for _, j := range indexList {
				if checkPaperRole(apenstaartje, onelineRunes, j) {
					countAdjecentRoles++
				}
			}
		}

		if countAdjecentRoles >= 4 {
			continue
		} else {
			freePaperRoles = append(freePaperRoles, i)
			totalFree += 1
		}

	}

	fmt.Println(freePaperRoles)
	//draw(freePaperRoles, onelineRunes)
	fmt.Println(totalFree, " = 13")
}

func day4Part2(filename string) {
	var freePaperRoles []int
	var totalFree int
	//var width int

	input := fileToArray(filename)

	var oneline string
	for _, line := range input {
		oneline += line
	}
	fmt.Println(oneline)
	onelineRunes := []rune(oneline)

	apenstaartje := onelineRunes[2]

	// loop as many times over the matrix as the total number of points
	for k := 0; k < len(onelineRunes)*137; k++ {
		// keep index in bound
		i := k % len(onelineRunes)

		if onelineRunes[i] != apenstaartje {
			continue
		}

		var countAdjecentRoles int

		indexList := [8]int{i - 137 - 1, i - 137, i - 137 + 1, i - 1, i + 1, i + 137 - 1, i + 137, i + 137 + 1}
		leftBorder := [5]int{i - 137, i - 137 + 1, i + 1, i + 137, i + 137 + 1}
		rightBorder := [5]int{i - 137 - 1, i - 137, i - 1, i + 137 - 1, i + 137}

		// if on left border
		if i%137 == 0 {
			for _, j := range leftBorder {
				if checkPaperRole(apenstaartje, onelineRunes, j) {
					countAdjecentRoles++
				}
			}
		}

		// if on right border
		if i%137 == 137-1 {
			for _, j := range rightBorder {
				if checkPaperRole(apenstaartje, onelineRunes, j) {
					countAdjecentRoles++
				}
			}
		}

		// Not on border
		if i%137 != 0 && i%137 != 137-1 {
			for _, j := range indexList {
				if checkPaperRole(apenstaartje, onelineRunes, j) {
					countAdjecentRoles++
				}
			}
		}

		if countAdjecentRoles >= 4 {
			continue
		} else {
			freePaperRoles = append(freePaperRoles, i)
			totalFree += 1
			onelineRunes[i] = '.'
		}
		if i%(18769) == 0 { // 137 * 137 = 18769
			fmt.Println(totalFree)
			//draw(freePaperRoles, onelineRunes)
		}

	}

	fmt.Println(freePaperRoles)
	//draw(freePaperRoles, onelineRunes)
	fmt.Println(totalFree, " = 13 of 43")
	fmt.Println("72297866 is to high")
}

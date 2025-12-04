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

func day4(filename string) {
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
		if i%137 == 9 {
			for _, j := range rightBorder {
				if checkPaperRole(apenstaartje, onelineRunes, j) {
					countAdjecentRoles++
				}
			}
		}

		// Not on border
		if i%137 != 0 && i%137 != 9 {
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
	fmt.Println(totalFree, " = 13")
}

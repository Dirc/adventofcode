package main

import "fmt"

func checkPaperRole(runes []rune, i int) bool {
	//var countAdjecentRoles int

	if onelineRunes[i] == apenstaartje {
		//fmt.Println("yes")
		countAdjecentRoles += 1
	}

	// check all cases
	i - 1
	i
	i + 1

	i - width - 1
	i - width
	i - width + 1

	i + width - 1
	i + width
	i + width + 1

	if i < 0 || i >= len(runes) {

	}
}

func day4(filename string) {
	var freePaperRoles []int
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
		fmt.Println(onelineRunes[i])

	}

	fmt.Println(len(freePaperRoles))
}

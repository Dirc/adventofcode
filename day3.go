package main

import (
	"fmt"
	"strconv"
)

func getHighest(list []int) (int, int) {
	highest := list[0]
	highestIndex := 0
	for i, item := range list {
		if item > highest {
			highest = item
			highestIndex = i
		}
	}
	return highest, highestIndex
}

func day3(filename string) {
	sum := 0
	input := fileToArray(filename)

	for _, item := range input {
		numericals := make([]int, len(item))
		for i, ch := range item {
			numericals[i], _ = strconv.Atoi(string(ch))
		}
		var joltageNumbers []int
		startIndex := 0
		for i := 1; i <= 12; i++ {
			highest, highestIndex := getHighest(numericals[startIndex : len(numericals)-12+i])
			joltageNumbers = append(joltageNumbers, highest)
			startIndex = startIndex + highestIndex + 1
		}
		// ex 1
		//highest, highestIndex := getHighest(numericals[:len(numericals)-1])
		//secondHighest, _ := getHighest(numericals[highestIndex+1:])

		joltage := joltageNumbers[0]*100000000000 + joltageNumbers[1]*10000000000 + joltageNumbers[2]*1000000000 + joltageNumbers[3]*100000000 + joltageNumbers[4]*10000000 + joltageNumbers[5]*1000000 + joltageNumbers[6]*100000 + joltageNumbers[7]*10000 + joltageNumbers[8]*1000 + joltageNumbers[9]*100 + joltageNumbers[10]*10 + joltageNumbers[11]

		fmt.Println(joltage)

		sum = sum + joltage
	}
	fmt.Println(sum)
}

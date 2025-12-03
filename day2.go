package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findDoubles(first string, last string) []string {
	var doubles []string
	firstInt, _ := strconv.Atoi(first)
	lastInt, _ := strconv.Atoi(last)
	for i := firstInt; i <= lastInt; i++ {
		iString := strconv.Itoa(i)
		// if lenth is odd, skip
		if len(iString)%2 == 1 {
			continue
		} else { // cut in half and compare
			half := len(iString) / 2
			if iString[half:] == iString[:half] {
				doubles = append(doubles, iString)
				fmt.Println(iString)
				continue
			}
			if repeatedSequence(iString) {
				doubles = append(doubles, iString)
				fmt.Println(iString)
			}
		}
	}
	return doubles
}

func repeatedSequence(number string) bool {
	// Check of 1e character repeated is, of de eerste 2 characters repeated zijn, enz.
	for i := 1; i <= len(number)/2; i++ {
		// Try rootNumbers from 1 character up to half the number
		rootNumber := number[0:i]
		// Check number in steps; len(rootNumber)
		stepSize := len(rootNumber)
		for j := 1; j < len(number)/stepSize; j++ {
			numberPart := number[j*stepSize : (j+1)*stepSize]
			if rootNumber != numberPart {
				break
			}
			// if last part of string/number reached: return true
			if (j+1)*stepSize == len(number) {
				print(number)
				print(" ")
				print("true")
				return true
			}
		}
	}
	print(number)
	print(" ")
	print("false")
	return false
}

func totalDoubles(filename string) []string {
	var doubles []string
	inputDay2 := fileToString(filename)

	input := strings.Split(inputDay2, ",")

	for _, item := range input {
		fmt.Println(item)
		idRange := strings.Split(item, "-")
		first := idRange[0]
		last := idRange[1]
		doubles = append(doubles, findDoubles(first, last)...)
	}
	return doubles

}

func day2(filename string) {
	doubles := totalDoubles(filename)
	var idSum = 0
	for _, id := range doubles {
		idInt, _ := strconv.Atoi(id)
		idSum = idSum + idInt
	}
	fmt.Println(idSum)
}

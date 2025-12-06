package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rmSpaces(input []string) []string {
	var output []string
	for _, item := range input {
		if item != "" {
			output = append(output, item)
		}
	}
	return output
}

func getInputInfo(filename string) {
	input := FileToArray(filename)

	// operators
	operatorsRaw := strings.Split(input[3], " ")
	operators := rmSpaces(operatorsRaw)

	// matrix
	columns := len(operators)
	rows := len(input) - 1

	fmt.Println(rows, columns)
}

func day6Part1(filename string) {
	input := FileToArray(filename)

	// operators
	operatorsRaw := strings.Split(input[4], " ")
	operators := rmSpaces(operatorsRaw)

	// matrix
	var matrix [4][1000]int
	columns := len(operators)
	rows := len(input) - 1

	for i, item := range input {
		if i == rows { // line with operators
			break
		}
		inputSplit := strings.Split(item, " ")
		inputPure := rmSpaces(inputSplit)

		for j, item := range inputPure {
			integer, _ := strconv.Atoi(item)
			matrix[i][j] = integer
		}
	}

	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println(operators)

	total := 0
	for i := 0; i < columns; i++ {
		columnTotal := 0
		for j := 0; j < rows; j++ {
			if operators[i] == "+" {
				columnTotal += matrix[j][i]
			}
			if operators[i] == "*" {
				if columnTotal == 0 {
					columnTotal = 1
				}
				columnTotal = columnTotal * matrix[j][i]
			}
		}
		fmt.Println(columnTotal)
		total += columnTotal
	}

	fmt.Println(total)
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func fileToArray(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return input

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func day1(filename string) {
	inputDay1 := fileToArray(filename)

	var dail = 50
	var count0 = 0

	for i, item := range inputDay1 {
		var rotation int
		var dail_new int
		rotation, err := strconv.Atoi(item[1:])
		if err != nil {
			panic(err)
		}

		if strings.HasPrefix(item, "L") {
			dail_new = dail - rotation
		} else {
			dail_new = dail + rotation
		}

		// ex. 1
		if dail_new == 0 {
			count0++
			dail = dail_new
		}

		// ex. 2
		//if abs(dail_new)+abs(dail) != abs(dail_new+dail) {
		//	count0++
		//}
		if dail_new < 0 {
			quotient := dail_new / (-100)
			count0 = count0 + 1 + quotient
			dail = dail_new + 100
		}
		if dail_new > 100 {
			quotient := dail_new / 100
			count0 = count0 + quotient
			dail = dail % 100
		}

		println(i, dail, count0)
	}
	print(count0)
}

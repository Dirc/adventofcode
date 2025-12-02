package main

import (
	"fmt"
	"strings"
)

func day2(filename string) {
	inputDay2 := fileToString(filename)

	input := strings.Split(inputDay2, ",")

	for _, item := range input {
		fmt.Println(item)
	}

}

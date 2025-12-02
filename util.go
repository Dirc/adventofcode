package main

import (
	"bufio"
	"os"
)

func fileToString(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return ""

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

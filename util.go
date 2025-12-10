package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FileToArray(filename string) []string {
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

func FileToString(filename string) string {
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Contains(intList []int, i int) bool {
	for _, v := range intList {
		if v == i {
			return true
		}
	}
	return false
}

func DrawStringArray(matrix []string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			print(matrix[i][j])
			if j == len(matrix)-1 {
				print("\n")
			}
		}
	}
}

func drawMatrix(matrix [][]string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			print(matrix[i][j])
			if j == len(matrix)-1 {
				print("\n")
			}
		}
	}
}

func Order(i, j int) (int, int) {
	if i < j {
		return i, j
	}
	return j, i
}

func GetHighest(filename string) {
	high0 := 0
	high1 := 0
	coordinates := File2Matrix(filename)

	for _, row := range coordinates {
		if row[0] > high0 {
			high0 = row[0]
		}
		if row[1] > high1 {
			high1 = row[1]
		}
	}
	fmt.Println(high0, high1)
}

func createFile(filename string, content string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func downloadInput(day string) {
	url := "https://adventofcode.com/2025/day/" + day + "/input"
	filename := "day" + day + ".txt"
	sessionCookie := os.Getenv("SESSION_COOKIE")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", "session="+sessionCookie)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("bad status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	createFile(filename, string(body))

}

func DailySetup(day string) {

	goFile := "day" + day + ".go"
	goFileContent := "package main\nimport \"fmt\"\nfunc day" + day + "part1(filename string) {\n\tfmt.Println(\"day" + day + "\")\n}"
	createFile(goFile, goFileContent)

	devFile := "day" + day + "-dev.txt"
	createFile(devFile, "")

	downloadInput(day)

}

package main

import (
	"strconv"
	"strings"
)

func day1(filename string) {
	inputDay1 := FileToArray(filename)

	var dail = 50
	var count0 = 0

	for i, item := range inputDay1 {
		var rotation int
		rotation, err := strconv.Atoi(item[1:])
		if err != nil {
			panic(err)
		}

		var dail_new int

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
			if dail != 0 {
				quotient := dail_new / (-100)
				count0 = count0 + 1 + quotient
				dail_new = dail_new % 100
				dail = 100 + dail_new // Note: dail_new is negative here
			}
			if dail == 0 {
				dail = 100 + dail_new
			}
		}
		if dail_new > 99 {
			quotient := dail_new / 100
			count0 = count0 + quotient
			dail = dail_new % 100
		}
		if dail_new > 0 && dail_new <= 99 {
			dail = dail_new
		}

		println(i, dail, item, count0)
	}
	print(count0)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {

	contents := readToSlice("../day1/sample.txt")

	// part 1

	var calibrationValues []string
	sumCalibrationValues := 0

	// loops through each item in the content array
	for i := 0; i < len(contents); i++ {
		word := contents[i]
		//fmt.Println(word)

		var firstDigit string
		var lastDigit string

		// loops through each char item at index (i) to check for first digit
		for j := 0; j < len(word); j++ {
			char := rune(word[j])
			if unicode.IsDigit(char) {
				firstDigit = string(char)
				break
			}

		}

		// loops backwards through each char item at index (i) to check for last digit
		for k := len(word) - 1; k >= 0; k-- {
			char := rune(word[k])
			if unicode.IsDigit(char) {
				lastDigit = string(char)
				break
			}
		}

		calibrationVal := firstDigit + lastDigit

		calibrationValues = append(calibrationValues, calibrationVal)

	}

	// sum of all calibration values
	for _, v := range calibrationValues {
		val, _ := strconv.Atoi(v)
		sumCalibrationValues += val
	}

	// part 1 result
	fmt.Println("Part 1 result: ", sumCalibrationValues)

}

func readToSlice(filePath string) []string {
	file, err := os.Open(filePath)
	checkErr(err)

	sc := bufio.NewScanner(file)

	var lines []string

	for sc.Scan() {
		line := sc.Text()

		lines = append(lines, line)

	}
	return lines
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := []string{"./input.txt"}

	file_data := parseFile(filename[:])

	sum := 0

	for i := 0; i < len(file_data); i++ {
		value := getCalibrateValue(file_data[i])
		sum += value
		fmt.Printf("%d: %s\n", value, file_data[i])
	}

	println(sum)
}

func getCalibrateValue(line string) int {
	first_digit := -1
	last_digit := -1

	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			if first_digit == -1 {
				first_digit = int(line[i]) - '0'
			} else {
				last_digit = int(line[i]) - '0'
			}
		}
	}

	if last_digit == -1 {
		last_digit = first_digit
	}

	return first_digit*10 + last_digit
}

func parseFile(files []string) []string {
	var rtnData []string

	for i := 0; i < len(files); i++ {
		f, err := os.Open(files[i])
		check(err)
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSuffix(line, "\n")
			rtnData = append(rtnData, line)
		}

		if scanner.Err() != nil {
			panic(scanner.Err())
		}
	}

	return rtnData
}

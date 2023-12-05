package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type ConversionRow struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

type ConversionTable struct {
	rows []ConversionRow
	name string
}

type Almanac struct {
	seeds            []int
	conversionTables []ConversionTable
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func returnMappedValue(source_val int, table ConversionTable) int {
	for i := 0; i < len(table.rows); i++ {
		start := table.rows[i].sourceStart
		rangeLen := table.rows[i].rangeLen
		if source_val >= start && source_val <= start+rangeLen {
			offset := source_val - start
			dest := table.rows[i].destStart
			return dest + offset
		}
	}

	return source_val
}

func getFinalValue(tables []ConversionTable, seed int) int {
	curr := seed
	for i := 0; i < len(tables); i++ {
		curr = returnMappedValue(curr, tables[i])
	}
	return curr
}

func main() {
	filename := "./input.txt"

	almanacData := parseFile(filename)

	fmt.Println(almanacData)

	lowestLocation := math.MaxInt64

	for i := 0; i < len(almanacData.seeds); i = i + 2 {
		seedStart := almanacData.seeds[i]
		seedRange := almanacData.seeds[i+1]
		for j := seedStart; j <= seedStart+seedRange; j++ { // lol this takes 2 min to run 💀
			val := getFinalValue(almanacData.conversionTables, j)
			if val < lowestLocation {
				lowestLocation = val
			}
		}
	}

	fmt.Printf("Lowest Location: %d\n", lowestLocation)
}

func parseFile(file string) Almanac {
	var rtnData Almanac
	rtnData.conversionTables = make([]ConversionTable, 0)
	rtnData.seeds = make([]int, 0)

	f, err := os.Open(file)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if scanner.Scan() {
		seedsLine := scanner.Text()
		re := regexp.MustCompile(`(\d+)`)
		match := re.FindAllString(seedsLine, -1)

		for i := 0; i < len(match); i++ {
			val, err := strconv.Atoi(match[i])
			check(err)
			rtnData.seeds = append(rtnData.seeds, val)
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue // skip empty lines
		}

		re := regexp.MustCompile(`(\d+)`)
		match := re.FindAllString(line, -1)

		if len(match) <= 0 { // title string means new table
			var newTable ConversionTable
			newTable.rows = make([]ConversionRow, 0)

			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					break
				}

				re := regexp.MustCompile(`(\d+)`)
				match := re.FindAllString(line, -1)

				match_ints := make([]int, 0)

				for i := 0; i < len(match); i++ {
					val, err := strconv.Atoi(match[i])
					check(err)
					match_ints = append(match_ints, val)
				}

				newRow := ConversionRow{match_ints[0], match_ints[1], match_ints[2]}

				newTable.rows = append(newTable.rows, newRow)
			}
			rtnData.conversionTables = append(rtnData.conversionTables, newTable)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rtnData
}

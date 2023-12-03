package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type part struct {
	number int
	isPart bool
	len    int
}

type engineParts struct {
	m_rows int
	m_cols int
	data   [][]*part
}

type engineGears struct {
	m_rows int
	m_cols int
	data   [][]*engineGear
}

type engineGear struct {
	parts []*part
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "./input.txt"

	schematic_parts := parseEngine(filename)
	fmt.Println("Engine parts parsed")
	schematic_gears := parseGears(filename)
	fmt.Println("Engine symbols parsed")
	gear_ratio_sum := 0

	for i := 0; i < schematic_gears.m_rows; i++ {
		for j := 0; j < schematic_gears.m_cols; j++ {
			if schematic_gears.data[i][j] != nil {
				getSurroundingParts(schematic_parts, schematic_gears, i, j)
			}
		}
	}

	for i := 0; i < schematic_gears.m_rows; i++ {
		for j := 0; j < schematic_gears.m_cols; j++ {
			if schematic_gears.data[i][j] != nil && len(schematic_gears.data[i][j].parts) == 2 {
				gear_ratio := schematic_gears.data[i][j].parts[0].number * schematic_gears.data[i][j].parts[1].number
				gear_ratio_sum += gear_ratio
			}
		}
	}

	fmt.Printf("Sum of gear ratio: %d\n", gear_ratio_sum)
}

func contains(s []*part, e *part) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func addToGear(gear *engineGear, part *part) {
	if !contains(gear.parts, part) {
		gear.parts = append(gear.parts, part)
	}
}

func getSurroundingParts(parts engineParts, gears engineGears, row int, col int) {
	if row-1 >= 0 && col-1 >= 0 && parts.data[row-1][col-1] != nil {
		addToGear(gears.data[row][col], parts.data[row-1][col-1])
	}

	if row-1 >= 0 && parts.data[row-1][col] != nil {
		addToGear(gears.data[row][col], parts.data[row-1][col])
	}

	if row-1 >= 0 && col+1 < parts.m_cols && parts.data[row-1][col+1] != nil {
		addToGear(gears.data[row][col], parts.data[row-1][col+1])
	}

	if col-1 >= 0 && parts.data[row][col-1] != nil {
		addToGear(gears.data[row][col], parts.data[row][col-1])
	}

	if col+1 < parts.m_cols && parts.data[row][col+1] != nil {
		addToGear(gears.data[row][col], parts.data[row][col+1])
	}

	if row+1 < parts.m_rows && col-1 >= 0 && parts.data[row+1][col-1] != nil {
		addToGear(gears.data[row][col], parts.data[row+1][col-1])
	}

	if row+1 < parts.m_rows && parts.data[row+1][col] != nil {
		addToGear(gears.data[row][col], parts.data[row+1][col])
	}

	if row+1 < parts.m_rows && col+1 < parts.m_cols && parts.data[row+1][col+1] != nil {
		addToGear(gears.data[row][col], parts.data[row+1][col+1])
	}
}

func parsePrefixInt(line string) (int, int) {
	var rtn int
	var length int

	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			rtn *= 10
			length++
			rtn += int(line[i]) - '0'
		} else {
			break
		}
	}

	return rtn, length
}

func parseEngine(file string) engineParts {
	var rtnData engineParts

	f, err := os.Open(file)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSuffix(line, "\n")
		rtnData.m_cols = len(line)
		rtnData.m_rows++

		rtnData.data = append(rtnData.data, make([]*part, rtnData.m_cols))

		for j := 0; j < rtnData.m_cols; j++ {
			if line[j] == '.' || isSymbol(rune(line[j])) {
				rtnData.data[i][j] = nil
			} else {
				value, length := parsePrefixInt(line[j:])
				new_part := part{value, false, length}
				fmt.Printf("Part: %d (%d,%d)-(%d,%d)\n", new_part.number, i, j, i, j+length-1)
				for k := j; k < j+length; k++ {
					rtnData.data[i][k] = &new_part
				}
				j += length - 1
			}
		}
		i++
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rtnData
}

func isSymbol(r rune) bool {
	if r == '.' || unicode.IsDigit(r) {
		return false
	}

	return true
}

func parseGears(file string) engineGears {
	var rtnData engineGears

	f, err := os.Open(file)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSuffix(line, "\n")
		rtnData.m_cols = len(line)
		rtnData.m_rows++

		rtnData.data = append(rtnData.data, make([]*engineGear, rtnData.m_cols))

		for j := 0; j < rtnData.m_cols; j++ {
			if rune(line[j]) == '*' {
				new_gear := engineGear{[]*part{}}
				rtnData.data[i][j] = &new_gear
			} else {
				rtnData.data[i][j] = nil
			}
		}
		i++
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rtnData
}

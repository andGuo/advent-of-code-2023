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

type engineSymbols struct {
	m_rows int
	m_cols int
	data   [][]rune
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
	schematic_symbols := parseSymbols(filename)
	fmt.Println("Engine symbols parsed")
	parts_sum := 0

	for i := 0; i < schematic_symbols.m_rows; i++ {
		for j := 0; j < schematic_symbols.m_cols; j++ {
			if isSymbol(schematic_symbols.data[i][j]) {
				fmt.Printf("Symbol: %c (%d,%d)\n", schematic_symbols.data[i][j], i, j)
				setSurroundingParts(schematic_parts, i, j)
			}
		}
	}

	for i := 0; i < schematic_parts.m_rows; i++ {
		for j := 0; j < schematic_parts.m_cols; j++ {
			if schematic_parts.data[i][j] != nil && schematic_parts.data[i][j].isPart {
				parts_sum += schematic_parts.data[i][j].number
				j += schematic_parts.data[i][j].len - 1
				fmt.Printf("Part: %d\n", schematic_parts.data[i][j].number)
			}
		}
	}

	fmt.Printf("Sum of parts: %d\n", parts_sum)
}

func setSurroundingParts(parts engineParts, row int, col int) {
	if row-1 >= 0 && col-1 >= 0 && parts.data[row-1][col-1] != nil {
		parts.data[row-1][col-1].isPart = true
	}

	if row-1 >= 0 && parts.data[row-1][col] != nil {
		parts.data[row-1][col].isPart = true
	}

	if row-1 >= 0 && col+1 < parts.m_cols && parts.data[row-1][col+1] != nil {
		parts.data[row-1][col+1].isPart = true
	}

	if col-1 >= 0 && parts.data[row][col-1] != nil {
		parts.data[row][col-1].isPart = true
	}

	if col+1 < parts.m_cols && parts.data[row][col+1] != nil {
		parts.data[row][col+1].isPart = true
	}

	if row+1 < parts.m_rows && col-1 >= 0 && parts.data[row+1][col-1] != nil {
		parts.data[row+1][col-1].isPart = true
	}

	if row+1 < parts.m_rows && parts.data[row+1][col] != nil {
		parts.data[row+1][col].isPart = true
	}

	if row+1 < parts.m_rows && col+1 < parts.m_cols && parts.data[row+1][col+1] != nil {
		parts.data[row+1][col+1].isPart = true
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

func parseSymbols(file string) engineSymbols {
	var rtnData engineSymbols

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

		rtnData.data = append(rtnData.data, make([]rune, rtnData.m_cols))

		for j := 0; j < rtnData.m_cols; j++ {
			if isSymbol(rune(line[j])) {
				rtnData.data[i][j] = rune(line[j])
			} else {
				rtnData.data[i][j] = '.'
			}
		}
		i++
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rtnData
}

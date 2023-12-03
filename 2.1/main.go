package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cube_set struct {
	red_cubes   int
	green_cubes int
	blue_cubes  int
}

type game struct {
	id   int
	sets []cube_set
}

type bag struct {
	red_cubes   int
	green_cubes int
	blue_cubes  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const total_red_cubes = 12
	const total_green_cubes = 13
	const total_blue_cubes = 14

	elf_bag := bag{total_red_cubes, total_green_cubes, total_blue_cubes}

	filename := []string{"./input.txt"}

	game_data := parseFile(filename[:])
	id_sum := 0

	for i := 0; i < len(game_data); i++ {
		if isPossibleGame(elf_bag, game_data[i]) {
			id_sum += game_data[i].id
		}
	}

	fmt.Printf("Sum of IDs: %d\n", id_sum)
}

func isValidSet(bag bag, set cube_set) bool {
	if bag.red_cubes-set.red_cubes >= 0 && bag.green_cubes-set.green_cubes >= 0 && bag.blue_cubes-set.blue_cubes >= 0 {
		return true
	}

	return false
}

func isPossibleGame(bag bag, game game) bool {
	for i := 0; i < len(game.sets); i++ {
		if !isValidSet(bag, game.sets[i]) {
			return false
		}
	}

	return true
}

func parseCubeValues(line string) cube_set {
	rtnData := cube_set{0, 0, 0}
	value_colour_strs := strings.Split(line, ", ")

	for i := 0; i < len(value_colour_strs); i++ {
		value, err := strconv.Atoi(strings.Split(value_colour_strs[i], " ")[0])

		if err != nil {
			panic(err)
		}

		switch strings.Split(value_colour_strs[i], " ")[1] {
		case "red":
			rtnData.red_cubes = value
		case "green":
			rtnData.green_cubes = value
		case "blue":
			rtnData.blue_cubes = value
		}
	}

	return rtnData
}

func parseCubeSets(line string) []cube_set {
	var rtnData []cube_set

	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimPrefix(line, " ")

	set_strings := strings.Split(line, ";")

	for i := 0; i < len(set_strings); i++ {
		temp_str := set_strings[i]
		temp_str = strings.TrimPrefix(temp_str, " ")
		new_set := parseCubeValues(temp_str)

		rtnData = append(rtnData, new_set)
	}

	return rtnData
}

func parseFile(files []string) []game {
	var rtnData []game

	for i := 0; i < len(files); i++ {
		f, err := os.Open(files[i])
		check(err)
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			var new_game game
			line := scanner.Text()
			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimPrefix(line, "Game ")

			new_game.id, err = strconv.Atoi(strings.Split(line, ":")[0])

			if err != nil {
				panic(err)
			}

			new_game.sets = parseCubeSets(strings.Split(line, ":")[1])

			rtnData = append(rtnData, new_game)
		}

		if scanner.Err() != nil {
			panic(scanner.Err())
		}
	}

	return rtnData
}

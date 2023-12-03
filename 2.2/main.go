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

type min_game struct {
	id  int
	bag bag
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
	filename := []string{"./input.txt"}

	game_data := parseFile(filename[:])
	power_sum := 0

	for i := 0; i < len(game_data); i++ {
		temp_bag := findMinBagSet(game_data[i])
		power_sum += setPowerValue(temp_bag)
	}

	fmt.Printf("Sum of Power of Sets: %d\n", power_sum)
}

func setPowerValue(min_game min_game) int {
	t_bag := min_game.bag
	return t_bag.red_cubes * t_bag.green_cubes * t_bag.blue_cubes
}

func findMinBagSet(game game) min_game {
	var rtnData min_game
	rtnData.id = game.id

	min_bag := bag{0, 0, 0}

	for i := 0; i < len(game.sets); i++ {
		if game.sets[i].red_cubes > min_bag.red_cubes {
			min_bag.red_cubes = game.sets[i].red_cubes
		}
		if game.sets[i].green_cubes > min_bag.green_cubes {
			min_bag.green_cubes = game.sets[i].green_cubes
		}
		if game.sets[i].blue_cubes > min_bag.blue_cubes {
			min_bag.blue_cubes = game.sets[i].blue_cubes
		}
	}

	rtnData.bag = min_bag

	return rtnData
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

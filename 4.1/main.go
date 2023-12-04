package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type winning_set struct {
	nums map[int]bool
}

type chosen_set struct {
	nums []int
}

type card struct {
	id          int
	winning     winning_set
	chosen      chosen_set
	num_winning int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "./input.txt"

	card_data := parseFile(filename)
	points_total := 0

	for i := 0; i < len(card_data); i++ {
		points_total += getCardPoints(&card_data[i])
		fmt.Printf("Card %d: %d (%d)\n", card_data[i].id, getCardPoints(&card_data[i]), card_data[i].num_winning)
	}

	fmt.Printf("Sum of Points: %d\n", points_total)
}

func getCardPoints(card *card) int {
	card.num_winning = 0

	for i := 0; i < len(card.chosen.nums); i++ {
		if card.winning.nums[card.chosen.nums[i]] {
			card.num_winning++
		}
	}

	score := 0

	if card.num_winning > 0 {
		score = int(math.Pow(2, float64(card.num_winning-1)))
	}

	return score
}

func parseCardLists(line string) (winning_set, chosen_set) {
	var winning winning_set
	var chosen chosen_set

	winning.nums = make(map[int]bool)
	chosen.nums = make([]int, 0)

	win_str := strings.Split(strings.Split(line, " | ")[0], " ")
	chosen_str := strings.Split(strings.Split(line, " | ")[1], " ")

	for i := 0; i < len(win_str); i++ {
		num_str := strings.ReplaceAll(win_str[i], " ", "")

		if num_str == "" {
			continue
		}

		num, err := strconv.Atoi(num_str)
		check(err)
		winning.nums[num] = true
	}

	for i := 0; i < len(chosen_str); i++ {
		num_str := strings.ReplaceAll(chosen_str[i], " ", "")

		if num_str == "" {
			continue
		}

		num, err := strconv.Atoi(num_str)
		check(err)
		chosen.nums = append(chosen.nums, num)
	}

	return winning, chosen
}

func parseFile(file string) []card {
	var rtnData []card

	f, err := os.Open(file)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var new_card card

		line := scanner.Text()
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimPrefix(line, "Card ")

		card_id_str := strings.Split(line, ": ")[0]
		card_id_str = strings.ReplaceAll(card_id_str, " ", "")
		new_card.id, err = strconv.Atoi(card_id_str)

		if err != nil {
			panic(err)
		}

		new_card.winning, new_card.chosen = parseCardLists(strings.Split(line, ": ")[1])

		rtnData = append(rtnData, new_card)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rtnData
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winning map[int64]bool
	chosen  map[int64]bool
}

func (c *Card) count() int64 {
	var count int64 = 0
	for win := range c.winning {
		if c.chosen[win] {
			count++
		}
	}
	return count
}

func (c *Card) score() int64 {
	cards := c.count()
	if cards > 0 {
		return 1 << (cards - 1)
	}
	return 0
}

type Runner struct {
	cards []Card
}

func (r *Runner) parse() {
	r.cards = make([]Card, 0)
	reader := bufio.NewReader(os.Stdin)
	for row := 0; ; row++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Reader Error: %s\n", err)
			return
		}

		data := strings.Split(line, ": ")
		nums := data[1]
		cards := strings.Split(nums, " | ")
		win, chose := cards[0], cards[1]

		winning := make(map[int64]bool)
		for _, stringValue := range strings.Fields(win) {
			value, err := strconv.Atoi(stringValue)
			if err != nil {
				fmt.Printf("Error converting string to int: %v\n", err)
				return
			}
			winning[int64(value)] = true
		}

		chosen := make(map[int64]bool)
		for _, stringValue := range strings.Fields(chose) {
			value, err := strconv.Atoi(stringValue)
			if err != nil {
				fmt.Printf("Error converting string to int: %v\n", err)
				return
			}
			chosen[int64(value)] = true
		}

		r.cards = append(r.cards, Card{winning, chosen})
	}
}

func (r *Runner) part1() {
	var total int64 = 0
	for _, card := range r.cards {
		total += card.score()
	}
	fmt.Printf("%v\n", total)
}

func (r *Runner) part2() {
	var total int64 = 0
	var mul = make([]int64, len(r.cards))
	var count int
	for i := range mul {
		mul[i] = 1
	}
	for i, card := range r.cards {
		count = int(card.count())
		for j := i + 1; j < i+count+1; j++ {
			mul[j] += mul[i]
		}
	}
	for _, value := range mul {
		total += value
	}
	fmt.Printf("%v\n", total)
}

func main() {
	r := Runner{}
	r.parse()
	r.part2()
}

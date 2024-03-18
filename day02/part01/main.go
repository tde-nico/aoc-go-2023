package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Turn struct {
	red   uint8
	green uint8
	blue  uint8
}

type Runner struct {
	gameList [][]Turn
}

func (t Turn) is_valid() bool {
	return t.red <= 12 && t.green <= 13 && t.blue <= 14
}

func (r *Runner) parse() {
	r.gameList = make([][]Turn, 0, 10)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Reader Error: %s\n", err)
			return
		}
		game := strings.Split(line, ": ")[1]
		turnList := make([]Turn, 0, 5)

		for _, t := range strings.Split(game, "; ") {
			var turn Turn

			for _, c := range strings.Split(t, ", ") {
				parts := strings.Split(c, " ")
				amountStr, color := parts[0], parts[1]
				amount, err := strconv.Atoi(amountStr)
				if err != nil {
					panic("Atoi Error")
				}
				switch color[0] {
				case 'r':
					turn.red = uint8(amount)
				case 'g':
					turn.green = uint8(amount)
				case 'b':
					turn.blue = uint8(amount)
				}
			}
			turnList = append(turnList, turn)
		}
		r.gameList = append(r.gameList, turnList)
	}
}

func (r *Runner) part1() {
	var validGames uint32
	for i, game := range r.gameList {
		flag := true
		for _, turn := range game {
			if !turn.is_valid() {
				flag = false
				break
			}
		}
		if flag {
			validGames += uint32(i) + 1
		}
	}
	fmt.Printf("%v\n", validGames)
}

func main() {
	r := Runner{}
	r.parse()
	r.part1()
}

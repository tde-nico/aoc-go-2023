package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInts(s string) []int64 {
	ret := make([]int64, 0)
	for _, stringValue := range strings.Fields(s) {
		value, err := strconv.Atoi(stringValue)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			break
		}
		ret = append(ret, int64(value))
	}
	return ret
}

type RaceInfo struct {
	time   int64
	record int64
}

func (r *RaceInfo) wins() int64 {
	var a float64 = -1
	var b float64 = float64(r.time)
	var c float64 = float64(-r.record)

	first := math.Ceil((-b + math.Sqrt(b*b-4.0*a*c)) / (2.0 * a))
	second := math.Floor((-b - math.Sqrt(b*b-4.0*a*c)) / (2.0 * a))

	if (first * (b - first)) == float64(r.record) {
		first += 1.0
	}
	if (second * (b - second)) == float64(r.record) {
		second -= 1.0
	}

	return int64(second-first) + 1
}

type Runner struct {
	races []RaceInfo
}

func (r *Runner) parse() {
	r.races = make([]RaceInfo, 0)

	lines := make([]string, 0, 100)
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
		lines = append(lines, line)
	}

	times := parseInts(strings.Split(lines[0], ": ")[1])
	records := parseInts(strings.Split(lines[1], ": ")[1])

	for i := 0; i < len(times); i++ {
		r.races = append(r.races, RaceInfo{times[i], records[i]})
	}
}

func (r *Runner) part1() {
	var total int64 = 1
	for _, race := range r.races {
		total *= race.wins()
	}
	fmt.Printf("%v\n", total)
}

func main() {
	r := Runner{}
	r.parse()
	r.part1()
}

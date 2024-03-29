package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int64
	end   int64
}

func (r *Range) contains(val int64) bool {
	return r.start <= val && val < r.end
}

type SingleMap struct {
	rng   Range
	delta int64
}

type Mapping struct {
	maps []SingleMap
}

func (m *Mapping) new() {
	m.maps = make([]SingleMap, 0)
}

func (m *Mapping) add_mapping(dst int64, src int64, len int64) {
	m.maps = append(m.maps, SingleMap{
		rng:   Range{start: src, end: src + len},
		delta: dst - src,
	})
}

func (m *Mapping) apply_map(val int64) int64 {
	for _, mp := range m.maps {
		if mp.rng.contains(val) {
			return val + mp.delta
		}
	}
	return val
}

type Runner struct {
	seeds   []int64
	mapping []Mapping
}

func (r *Runner) parse() {
	r.seeds = make([]int64, 0)
	r.mapping = make([]Mapping, 0)

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

	seeds := strings.Split(lines[0], ": ")[1]
	for _, stringValue := range strings.Fields(seeds) {
		value, err := strconv.Atoi(stringValue)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return
		}
		r.seeds = append(r.seeds, int64(value))
	}

	var curmap Mapping
	curmap.new()
	for _, line := range lines[2:] {
		if len(line) == 1 {
			continue
		}
		if strings.Contains(line, ":") {
			r.mapping = append(r.mapping, curmap)
			curmap.new()
			continue
		}
		nums := make([]int64, 0)
		for _, stringValue := range strings.Fields(line) {
			value, err := strconv.Atoi(stringValue)
			if err != nil {
				fmt.Printf("Error converting string to int: %v\n", err)
				return
			}
			nums = append(nums, int64(value))
		}
		curmap.add_mapping(nums[0], nums[1], nums[2])
	}
	if len(curmap.maps) > 0 {
		r.mapping = append(r.mapping, curmap)
	}
}

func (r *Runner) part1() {
	var mn int64 = math.MaxInt64
	for _, seed := range r.seeds {
		cur := seed
		for _, mp := range r.mapping {
			cur = mp.apply_map(cur)
		}
		mn = min(mn, cur)
	}
	fmt.Printf("%v\n", mn)
}

func main() {
	r := Runner{}
	r.parse()
	r.part1()
}

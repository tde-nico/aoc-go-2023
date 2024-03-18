package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	y int64
	x int64
}

type PartNumber struct {
	value  int64
	points map[Point]bool
}

func (p *PartNumber) new(row int64, col int64, ch int64) {
	p.points = make(map[Point]bool)
	points := []Point{
		{row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1},
		{row, col - 1}, {row, col + 1},
		{row + 1, col - 1}, {row + 1, col}, {row + 1, col + 1},
	}
	for _, point := range points {
		p.points[point] = true
	}
	p.value = ch - 0x30
}

func (p *PartNumber) add_digit(row int64, col int64, ch int64) {
	points := []Point{
		{row - 1, col + 1},
		{row, col + 1},
		{row + 1, col + 1},
	}
	for _, point := range points {
		p.points[point] = true
	}
	p.value = p.value*10 + ch - 0x30
}

type Runner struct {
	nums  []PartNumber
	syms  map[Point]bool
	gears map[Point]bool
}

func (r *Runner) parse() {
	r.syms = make(map[Point]bool)
	r.gears = make(map[Point]bool)
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

		var cur_num PartNumber
		initialized := false
		for col, ch := range line {
			if ch >= 0x30 && ch <= 0x39 {
				if initialized {
					cur_num.add_digit(int64(row), int64(col), int64(ch))
				} else {
					initialized = true
					cur_num.new(int64(row), int64(col), int64(ch))
				}
			} else {
				if initialized {
					r.nums = append(r.nums, cur_num)
				}
				if ch != '.' && ch != '\n' {
					r.syms[Point{int64(row), int64(col)}] = true
					if ch == '*' {
						r.gears[Point{int64(row), int64(col)}] = true
					}
				}
				initialized = false
				cur_num = PartNumber{}
			}
		}
	}
}

func (r *Runner) part1() {
	total := 0
	for _, num := range r.nums {
		for point := range num.points {
			if r.syms[point] {
				total += int(num.value)
				break
			}
		}
	}
	fmt.Printf("%v\n", total)
}

func (r *Runner) part2() {
	total := 0
	for gear := range r.gears {
		matches := make([]int64, 0, 2)
		for _, num := range r.nums {
			if num.points[gear] {
				matches = append(matches, num.value)
			}
		}
		if len(matches) == 2 {
			total += int(matches[0]) * int(matches[1])
		}
	}
	fmt.Printf("%v\n", total)
}

func main() {
	r := Runner{}
	r.parse()
	r.part2()
}

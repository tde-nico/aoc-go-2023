package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i int
	var j int
	var left rune
	var right rune
	var sum int32

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
		i = 0
		for {
			if line[i] >= 0x30 && line[i] <= 0x39 {
				left = rune(line[i]) - 0x30
				break
			}
			i++
		}
		j = len(line) - 1
		for {
			if line[j] >= 0x30 && line[j] <= 0x39 {
				right = rune(line[j]) - 0x30
				break
			}
			j--
		}

		sum += left*10 + right
	}

	fmt.Printf("%v\n", sum)
}

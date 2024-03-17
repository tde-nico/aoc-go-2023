package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var left rune
	var right rune
	var sum int32
	var flag bool

	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		// Read line
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Reader Error: %s\n", err)
			return
		}

		// Parse left number
		for i := 0; i < len(line); i++ {
			flag = false
			for digit, value := range digits {
				fin := min(len(line)-1, i+len(digit))
				if line[i:fin] == digit {
					left = rune(value)
					flag = true
					break
				}
			}
			if flag {
				break
			}
			if line[i] >= 0x30 && line[i] <= 0x39 {
				left = rune(line[i]) - 0x30
				break
			}
		}

		// Parse right number
		for j := len(line) - 1; j >= 0; j-- {
			flag = false
			for digit, value := range digits {
				fin := min(len(line)-1, j+len(digit))
				if line[j:fin] == digit {
					right = rune(value)
					flag = true
					break
				}
			}
			if flag {
				break
			}
			if line[j] >= 0x30 && line[j] <= 0x39 {
				right = rune(line[j]) - 0x30
				break
			}
		}

		// Add to sum
		sum += left*10 + right
	}

	fmt.Printf("%v\n", sum)
}

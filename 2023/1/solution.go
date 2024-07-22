package main

import (
	"fmt"
	"os"
	"strings"
)

func gimmeFirstDigit(line string) int {
	first := -1

	for _, ch := range line {
		if ch >= '0' && ch <= '9' {
			first = int(ch - '0')
			break
		}
	}

	return first
}

func gimmeLastDigit(line string) int {
	sec := -1

	for ii := len(line) - 1; ii >= 0; ii-- {
		ch := line[ii]
		if ch >= '0' && ch <= '9' {
			sec = int(ch - '0')
			break
		}
	}

	return sec
}

func gimmeValFromLine(line string) int {
	return gimmeFirstDigit(line)*10 + gimmeLastDigit(line)
}

func firstSolution(file []byte) int {
	sum := 0

	for _, line := range strings.Split(string(file), "\n") {
		if len(line) == 0 {
			continue
		}

		sum += gimmeValFromLine(line)
	}

	return sum
}

var digitMaps = map[string]int{
	"zero":  0,
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

func gimmeFirstNum(line string) int {
	first := -1

	for idx, ch := range line {
		if ch >= '0' && ch <= '9' {
			first = int(ch - '0')
			break
		}

		for key := range digitMaps {
			if strings.HasPrefix(line[idx:], key) {
				return digitMaps[key]
			}
		}
	}

	return first
}

func gimmeLastNum(line string) int {
	sec := -1

	for idx := len(line) - 1; idx >= 0; idx-- {
		ch := line[idx]
		if ch >= '0' && ch <= '9' {
			sec = int(ch - '0')
			break
		}

		for key := range digitMaps {
			if strings.HasPrefix(line[idx:], key) {
				return digitMaps[key]
			}
		}

	}

	return sec
}

func secondSolution(file []byte) int {
	sum := 0
	for _, line := range strings.Split(string(file), "\n") {
		if len(line) == 0 {
			continue
		}

		first := gimmeFirstNum(line)
		second := gimmeLastNum(line)

		sum += first*10 + second
	}

	return sum

}

func main() {
	file, e := os.ReadFile("input.txt")
	if e != nil {
		panic(e)
	}

	fmt.Println("First solution is:\t", firstSolution(file))
	fmt.Println("Second solution is:\t", secondSolution(file))
}

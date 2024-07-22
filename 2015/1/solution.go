package main

import (
	"fmt"
	"os"
)

func main() {
	file, e := os.ReadFile("input.txt")
	if e != nil {
		panic(e)
	}

	floor := 0
	first := true

	for pos, ch := range file {
		if ch == '(' {
			floor += 1
		} else if ch == ')' {
			floor -= 1
		}

		if floor < 0 && first {
			first = false
			fmt.Println("Reached basement on pos:\t", pos+1)
		}
	}
	fmt.Println("Santa's floor should be:\t", floor)
}

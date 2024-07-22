package main

import (
	"fmt"
	"os"
	"strconv"
)

var delivered map[string]int

func deliverPresent(loc [2]int) {
	key := strconv.Itoa(loc[0]) + "," + strconv.Itoa(loc[1])

	delivered[key] += 1
}

func partOne() {
	file, e := os.ReadFile("input.txt")
	if e != nil {
		panic(e)
	}

	delivered = make(map[string]int)
	location := [2]int{0, 0}
	deliverPresent(location)

	for _, ch := range file {
		switch ch {
		case '^':
			location[1]++
			deliverPresent(location)
		case '>':
			location[0]++
			deliverPresent(location)
		case 'v':
			location[1]--
			deliverPresent(location)
		case '<':
			location[0]--
			deliverPresent(location)
		default:
			/* Do not deliver */
			break
		}
	}
}

func partTwo() {
	file, e := os.ReadFile("input.txt")
	if e != nil {
		panic(e)
	}

	delivered = make(map[string]int)
	santaLoc := [2]int{0, 0}
	deliverPresent(santaLoc)
	robosantaLoc := [2]int{0, 0}
	deliverPresent(robosantaLoc)

	santaTurn := true

	for _, ch := range file {
		if santaTurn {
			switch ch {
			case '^':
				santaLoc[1]++
				deliverPresent(santaLoc)
			case '>':
				santaLoc[0]++
				deliverPresent(santaLoc)
			case 'v':
				santaLoc[1]--
				deliverPresent(santaLoc)
			case '<':
				santaLoc[0]--
				deliverPresent(santaLoc)
			default:
				/* Do not deliver */
				break
			}
			santaTurn = false
		} else {
			switch ch {
			case '^':
				robosantaLoc[1]++
				deliverPresent(robosantaLoc)
			case '>':
				robosantaLoc[0]++
				deliverPresent(robosantaLoc)
			case 'v':
				robosantaLoc[1]--
				deliverPresent(robosantaLoc)
			case '<':
				robosantaLoc[0]--
				deliverPresent(robosantaLoc)
			default:
				/* Do not deliver */
				break
			}
			santaTurn = true
		}
	}
}

func main() {
	partOne()
	fmt.Println("Part 1 Total Houses: ", len(delivered))
	partTwo()
	fmt.Println("Part 2 Total Houses: ", len(delivered))
}

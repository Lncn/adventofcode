package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDims(dim string) []int {
	sides := make([]int, 3)

	lwh := strings.Split(dim, "x")
	if len(lwh) != 3 {
		return sides
	}

	var err error = nil
	sides[0], err = strconv.Atoi(lwh[0])
	if err != nil {
		return sides
	}
	sides[1], err = strconv.Atoi(lwh[1])
	if err != nil {
		return sides
	}
	sides[2], err = strconv.Atoi(lwh[2])
	if err != nil {
		return sides
	}

	return sides
}

func findSides(dim []int) []int {
	return []int{dim[0] * dim[1], dim[1] * dim[2], dim[0] * dim[2]}
}

func ribbonSize(sides []int) int {
	volume := sides[0] * sides[1] * sides[2]

	smallest := 0
	for idx, side := range sides {
		if side < sides[smallest] {
			smallest = idx
		}
	}

	if smallest == 0 {
		if sides[1] < sides[2] {
			return 2*sides[smallest] + 2*sides[1] + volume
		} else {
			return 2*sides[smallest] + 2*sides[2] + volume
		}
	} else if smallest == 1 {
		if sides[0] < sides[2] {
			return 2*sides[smallest] + 2*sides[0] + volume
		} else {
			return 2*sides[smallest] + 2*sides[2] + volume
		}
	} else {
		if sides[0] < sides[1] {
			return 2*sides[smallest] + 2*sides[0] + volume
		} else {
			return 2*sides[smallest] + 2*sides[1] + volume
		}
	}
}

func wrapSize(sides []int) int {
	smallest := sides[0]
	for _, side := range sides {
		if side < smallest {
			smallest = side
		}
	}

	return 2*sides[0] + 2*sides[1] + 2*sides[2] + smallest
}

func main() {
	file, e := os.Open("input.txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	papersum := 0
	ribbonsum := 0

	/* By default, Scan() splite on newline */
	for scanner.Scan() {
		dims := findDims(scanner.Text())
		sides := findSides(dims)
		ribbonsum += ribbonSize(dims)
		papersum += wrapSize(sides)
	}

	fmt.Println("Total wrapping paper:\t", papersum)
	fmt.Println("Total ribbon length:\t", ribbonsum)
}

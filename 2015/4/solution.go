package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func testHash5(hash [16]byte) bool {
	// We check first three bytes. Looking for 5 nibbles to be zero so...
	// hash = 00 00 0X XX..
	if hash[0] != 0 || hash[1] != 0 || hash[2]&0xF0 != 0 {
		return false
	}

	return true
}

func testHash6(hash [16]byte) bool {
	// We check first three bytes. Looking for 6 nibbles to be zero so...
	// hash = 00 00 00 XX..
	if hash[0] != 0 || hash[1] != 0 || hash[2] != 0 {
		return false
	}

	return true
}

func testString(key string, test int, signal chan int) {
	hash := md5.Sum([]byte(key + strconv.Itoa(test)))
	if testHash6(hash) {
		signal <- test
	}
}

func main() {
	// The example 'abcdef' should md5 hash with int '609043' to produce
	// five leading zeros: md5sum('abcdef609043') = 000001dbbfa...
	//key := "abcdef"
	//key := "pqrstuv" should find '1048970'
	key := "bgvyzdsv"

	/* Start with 1... */
	test := 1

	fmt.Println("Part 1:")
	for {
		hash := md5.Sum([]byte(key + strconv.Itoa(test)))
		if testHash5(hash) {
			fmt.Printf("Found MD5(%s%s) = %v\n", key, strconv.Itoa(test), hash)
			break
		}
		test += 1
	}

	/* Start with 1... */
	test = 1

	fmt.Println("Part 2:")
	for {
		hash := md5.Sum([]byte(key + strconv.Itoa(test)))
		if testHash6(hash) {
			fmt.Printf("Found MD5(%s%s) = %v\n", key, strconv.Itoa(test), hash)
			break
		}
		test += 1
	}

	fmt.Println("Just for fun... Go routine Part 2:")

	signal := make(chan int)
	go func() {
		test := 1
		for /* ever */ {
			go testString(key, test, signal)
			test += 1
		}
	}()

	/* This should block until a goroutine finds an winner! */
	winner := <-signal
	fmt.Printf("Found MD5 winner: %v\n", winner)
}

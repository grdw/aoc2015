package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const INPUT string = "ckczppom"

func main() {
	fmt.Println("Part 1:", p1(INPUT, "00000"))
	fmt.Println("Part 2:", p1(INPUT, "000000"))
}

func p1(input string, end string) int {
	start := 1

	for {
		comb := fmt.Sprintf("%s%d", input, start)
		result := md5.Sum([]byte(comb))

		if strings.HasPrefix(fmt.Sprintf("%x", result), end) {
			return start
		}

		start += 1
	}
}

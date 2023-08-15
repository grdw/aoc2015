package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const INPUT string = "ckczppom"

func main() {
	fmt.Println("Part 1:", calculateMd5(INPUT, "00000"))
	fmt.Println("Part 2:", calculateMd5(INPUT, "000000"))
}

func calculateMd5(input string, prefix string) int {
	start := 1

	for {
		comb := fmt.Sprintf("%s%d", input, start)
		result := md5.Sum([]byte(comb))

		if strings.HasPrefix(fmt.Sprintf("%x", result), prefix) {
			return start
		}

		start += 1
	}
}

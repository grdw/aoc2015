package main

import (
	"fmt"
)

const input int = 36000000

func main() {
	p1 := gift()
	fmt.Println("part 1:", p1)
}

func gift() int {
	houseIndex := 0
	incr := 10000

	for {
		total := presents(houseIndex)

		if total > input {
			if incr == 1 {
				return houseIndex
			}
			houseIndex -= incr * 100
			incr /= 10

		}
		houseIndex += incr
	}
}

func presents(i int) int {
	total := 0
	for j := 1; j <= i; j++ {
		if i%j == 0 {
			total += j * 10
		}
	}
	return total
}

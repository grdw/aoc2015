package main

import (
	"fmt"
)

const input int = 36000000

func main() {
	p1 := gift()
	fmt.Println("part 1:", p1)
	p2 := giftWithLimit()
	fmt.Println("part 2:", p2)
}

func gift() int {
	houseIndex := 0
	incr := 100

	for {
		total := presents(houseIndex)

		if total > input {
			return houseIndex
		}
		houseIndex += incr
	}
}

// TODO: This can be done better I guess...
// .. but oh whelp, not a math genius... so brute force it is!
func giftWithLimit() int {
	houseIndex := 0
	incr := 10

	for {
		total := presentsWithLimit(houseIndex)

		if total > input {
			return houseIndex
		}
		houseIndex += incr
	}
}

func presentsWithLimit(i int) int {
	total := 0
	limit := 50
	for j := 1; j <= i; j++ {
		if i%j == 0 && j*limit >= i {
			total += j * 11
		}
	}
	return total
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

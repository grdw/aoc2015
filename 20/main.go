package main

import (
	"fmt"
	"math"
)

const input int = 36000000

func main() {
	p1 := gift(10, math.MaxInt64)
	fmt.Println("part 1:", p1)
	p2 := gift(11, 50)
	fmt.Println("part 2:", p2)
}

func gift(s int, max int) int {
	size := input / 10
	var x [input / 10]int

	for i := 1; i < size; i++ {
		counter := 0
		for j := i; j < size-1; j += i {
			x[j] += s * i
			counter++
			if counter == max {
				break
			}
		}
	}

	fi := -1
	for i, v := range x {
		if v >= input {
			fi = i
			break
		}
	}

	return fi
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	directions, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Amount of houses that get at least 1 present: %d\n", houses(directions))
}

type point struct {
	x int
	y int
}

func houses(directions []byte) int {
	start := point{0, 0}
	counts := make(map[point]int)
	counts[start] += 1

	for _, d := range directions {
		switch d {
		case '<':
			start.x -= 1
			break
		case '>':
			start.x += 1
			break
		case 'v':
			start.y -= 1
			break
		case '^':
			start.y += 1
			break
		}

		counts[start] += 1
	}

	return len(counts)
}

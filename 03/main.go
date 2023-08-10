package main

import (
	"fmt"
	"log"
	"os"
)

type point struct {
	x int
	y int
}

func main() {
	directions, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Amount of houses that get at least 1 present: %d\n", routeWithSanta(directions))
	fmt.Printf("Amount of houses with Robo Santa that get at least 1 present: %d\n", routeWithRobo(directions))
}

func routeWithSanta(directions []byte) int {
	return len(route(directions))
}

func routeWithRobo(directions []byte) int {
	dirsSanta := []byte{}
	dirsRobo := []byte{}

	for i, d := range directions {
		if i%2 == 0 {
			dirsSanta = append(dirsSanta, d)
		} else {
			dirsRobo = append(dirsRobo, d)
		}
	}

	santaMap := route(dirsSanta)
	for k, v := range route(dirsRobo) {
		santaMap[k] += v
	}
	return len(santaMap)
}

func route(directions []byte) map[point]int {
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

	return counts
}

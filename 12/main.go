package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", parseInput(content))
}

func parseInput(text []byte) int {
	total := 0
	blurb := []byte{}
	for _, b := range text {
		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			blurb = append(blurb, b)
			break
		default:
			i, _ := strconv.Atoi(string(blurb))
			total += i
			blurb = []byte{}
			break
		}
	}
	return total
}

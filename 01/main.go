package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer to part 1 is: %d\n", countFloors(content))
	fmt.Printf("The answer to part 2 is: %d\n", basement(content))
}

func countFloors(content []byte) int {
	n := bytes.Count(content, []byte("("))
	m := bytes.Count(content, []byte(")"))
	return n - m
}

func basement(content []byte) int {
	floor := 0
	for i, b := range content {
		if b == '(' {
			floor += 1
		} else if b == ')' {
			floor -= 1
		}

		if floor == -1 {
			return i + 1
		}
	}
	return 0
}

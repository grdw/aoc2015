package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	strings := parse("input")
	fmt.Println("Part 1:", part1(strings))
	fmt.Println("Part 2:", part2(strings))
}

func part1(s []string) int {
	left := 0
	right := 0

	for _, s := range s {
		bs, _ := strconv.Unquote(s)
		left += len(s)
		right += len(bs)
	}

	return left - right
}

func part2(s []string) int {
	left := 0
	right := 0

	for _, s := range s {
		ss := strconv.Quote(s)
		left += len(ss)
		right += len(s)
	}

	return left - right
}

func parse(file string) []string {
	s := []string{}
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		s = append(s, line)
	}
	return s
}

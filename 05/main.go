package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	list, err := parse("input", matches)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", len(list))

	list, err = parse("input", improvedMatches)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2:", len(list))
}

func parse(file string, matcher func(line string) bool) ([]string, error) {
	st := []string{}
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		return st, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if matcher(line) {
			st = append(st, line)
		}
	}

	return st, nil
}

// Part one:
func matches(line string) bool {
	return vowelCountThree(line) && doubleMatch(line) && mismatch(line)
}

func doubleMatch(line string) bool {
	prevR := ' '
	for _, l := range line {
		if l == prevR {
			return true
		}
		prevR = l
	}
	return false
}

func vowelCountThree(line string) bool {
	count := 0
	for _, l := range line {
		switch l {
		case 'a', 'e', 'i', 'o', 'u':
			count += 1
			break
		}
	}

	return count > 2
}

func mismatch(line string) bool {
	mismatches := []string{"ab", "cd", "pq", "xy"}
	for _, m := range mismatches {
		if strings.Contains(line, m) {
			return false
		}
	}
	return true
}

// Part two:
func improvedMatches(line string) bool {
	return twins(line) && repeatBetween(line)
}

func twins(line string) bool {
	l := len(line)

	for i := 0; i < l-1; i++ {
		left := fmt.Sprintf("%c%c", line[i], line[i+1])
		if strings.Count(line, left) > 1 {
			return true
		}
	}

	return false
}

func repeatBetween(line string) bool {
	l := len(line)
	for i := 0; i < l-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

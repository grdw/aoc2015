package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type sue = map[string]int

func main() {
	auntieSues := parse("input")
	query := make(map[string]int)
	query["children"] = 3
	query["cats"] = 7
	query["samoyeds"] = 2
	query["pomeranians"] = 3
	query["akitas"] = 0
	query["vizslas"] = 0
	query["goldfish"] = 5
	query["trees"] = 3
	query["cars"] = 2
	query["perfumes"] = 1

	fmt.Println("Part 1:", findAuntieExact(auntieSues, query))
	fmt.Println("Part 2:", findAuntieWithRanges(auntieSues, query))
}

func findAuntieExact(auntieSues []sue, q map[string]int) int {
	for i, a := range auntieSues {
		match := true
		for k, v := range a {
			if q[k] != v {
				match = false
			}
		}

		if match {
			return i + 1
		}
	}
	return -1
}

func findAuntieWithRanges(auntieSues []sue, q map[string]int) int {
	for i, a := range auntieSues {
		match := true
		for k, v := range a {
			switch k {
			case "trees", "cats":
				if q[k] >= v {
					match = false
				}
				break
			case "pomeranians", "goldfish":
				if q[k] <= v {
					match = false
				}
				break
			default:
				if q[k] != v {
					match = false
				}
				break
			}
		}

		if match {
			return i + 1
		}
	}
	return -1
}

func parse(file string) []sue {
	sues := []sue{}
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		aunt := make(sue)
		re := regexp.MustCompile("[:, ]+")
		s := re.Split(line, -1)

		for i := 0; i < len(s); i += 2 {
			v, _ := strconv.Atoi(s[i+1])
			if s[i] == "Sue" {
				continue
			}
			aunt[s[i]] = v
		}
		sues = append(sues, aunt)
	}

	return sues
}

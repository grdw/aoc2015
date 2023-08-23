package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nameIndex int = 0
const speedIndex int = 3
const timeIndex int = 6
const pauseIndex int = 13

type reindeer struct {
	name  string
	speed int // in km/s
	time  int // in seconds
	pause int // in seconds
}

func main() {
	exMaxD, exMaxP := maxes(parse("input_test"), 1000)
	part1, part2 := maxes(parse("input"), 2503)
	fmt.Println("Ex:", exMaxD)
	fmt.Println("Part 1:", part1)
	fmt.Println("Ex:", exMaxP)
	fmt.Println("Part 2:", part2)
}

func maxes(rs []reindeer, seconds int) (int, int) {
	points := make(map[string]int)
	distance := make(map[string]int)
	cycles := make(map[string]int)
	clock := 0

	for clock <= seconds {
		for _, r := range rs {
			pausing := cycles[r.name] == r.time
			if !pausing {
				distance[r.name] += r.speed
				cycles[r.name] += 1
			} else if clock%(r.time+r.pause) == 0 {
				cycles[r.name] = 0
			}
		}

		winner, _ := max(distance)
		points[winner]++
		clock++
	}

	_, maxD := max(distance)
	_, maxP := max(points)
	return maxD, maxP
}

func max(m map[string]int) (string, int) {
	maxP := 0
	maxKey := ""
	for k, p := range m {
		if p > maxP {
			maxP = p
			maxKey = k
		}
	}
	return maxKey, maxP
}

func parse(file string) []reindeer {
	rs := []reindeer{}

	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		words := strings.Split(line, " ")

		speed, _ := strconv.Atoi(words[speedIndex])
		time, _ := strconv.Atoi(words[timeIndex])
		pause, _ := strconv.Atoi(words[pauseIndex])
		rs = append(rs, reindeer{words[nameIndex], speed, time, pause})
	}

	return rs
}

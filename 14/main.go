package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type reindeer struct {
	name  string
	speed int // in km/s
	time  int // in seconds
	pause int // in seconds
}

func main() {
	fmt.Println("Ex:", travel(parse("input_test"), 1000))
	fmt.Println("Part 1:", travel(parse("input"), 2503))
}

func travel(rs []reindeer, seconds int) int {
	distance := make(map[string]int)

	for _, r := range rs {
		clock := 0
		cycle := 0

		for clock < seconds {
			p := cycle % r.time
			distance[r.name] += r.speed

			if p == (r.time - 1) {
				clock += (r.pause + 1)
				cycle = 0
			} else {
				cycle++
				clock++
			}
		}
	}

	fmt.Println(distance)
	max := 0
	for _, d := range distance {
		if d > max {
			max = d
		}
	}

	return max
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

		speed, _ := strconv.Atoi(words[3])
		time, _ := strconv.Atoi(words[6])
		pause, _ := strconv.Atoi(words[13])
		rs = append(rs, reindeer{words[0], speed, time, pause})
	}

	return rs
}

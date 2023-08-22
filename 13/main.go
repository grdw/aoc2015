package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type calcHeapFunc = func(a []string, hi map[string]int) int

func main() {
	guests, happinessIndex := parse("input")
	max := 0
	heapPermutation(
		guests,
		len(guests),
		happinessIndex,
		calculateHappiness,
		&max,
	)
	fmt.Println("Part 1:", max)
}

func calculateHappiness(guests []string, happinessIndex map[string]int) int {
	total := 0
	for i := 0; i < len(guests); i++ {
		nextIndex := (i + 1) % len(guests)
		lr := fmt.Sprintf("%s-%s", guests[i], guests[nextIndex])
		rl := fmt.Sprintf("%s-%s", guests[nextIndex], guests[i])
		total = total + happinessIndex[lr] + happinessIndex[rl]
	}
	return total
}

func heapPermutation(
	a []string,
	size int,
	happinessIndex map[string]int,
	calc calcHeapFunc,
	max *int,
) {
	if size == 1 {
		newMax := calc(a, happinessIndex)
		if newMax > *max {
			*max = newMax
		}
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1, happinessIndex, calc, max)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func parse(file string) ([]string, map[string]int) {
	guests := []string{}
	guestMap := make(map[string]bool)
	happinessIndex := make(map[string]int)
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		re := regexp.MustCompile(
			`([A-Za-z]+) would (lose|gain) (\d+) happiness units by sitting next to ([A-Za-z]+)`,
		)

		matches := re.FindAllSubmatch(line, -1)
		p1 := string(matches[0][1])
		lossGain := string(matches[0][2])
		hap, _ := strconv.Atoi(string(matches[0][3]))
		p2 := string(matches[0][4])
		if lossGain == "lose" {
			hap *= -1
		}
		happinessIndex[fmt.Sprintf("%s-%s", p1, p2)] = hap
		guestMap[p1] = true
	}

	for guest, _ := range guestMap {
		guests = append(guests, guest)
	}

	return guests, happinessIndex
}

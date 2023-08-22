package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type calcHeapFunc = func(a []string, hi map[string]int) int

func main() {
	guests, happinessIndex := parse("input")
	fmt.Println("Part 1:", maxHappiness(guests, happinessIndex))

	me := "me"
	for _, guest := range guests {
		happinessIndex[fmt.Sprintf("%s-%s", guest, me)] = 0
		happinessIndex[fmt.Sprintf("%s-%s", me, guest)] = 0
	}
	guests = append(guests, me)
	fmt.Println("Part 2:", maxHappiness(guests, happinessIndex))
}

func maxHappiness(guests []string, happinessIndex map[string]int) int {
	max := 0
	heapPermutation(
		guests,
		len(guests),
		happinessIndex,
		calculateHappiness,
		&max,
	)
	return max
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
	happinessIndex := make(map[string]int)

	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		words := strings.Split(line[0:len(line)-1], " ")

		p1 := words[0]
		p2 := words[len(words)-1]
		hap, _ := strconv.Atoi(words[3])
		if words[2] == "lose" {
			hap *= -1
		}

		happinessIndex[fmt.Sprintf("%s-%s", p1, p2)] = hap
		if !slices.Contains(guests, p1) {
			guests = append(guests, p1)
		}
	}

	return guests, happinessIndex
}

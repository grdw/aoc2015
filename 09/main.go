package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type edgeMap = map[string]int

func main() {
	edges, nodes := parse("input")
	min := math.MaxInt32
	max := 0
	maxPerms := factorial(len(nodes))

	for maxPerms > 1 {
		d := distance(nodes, edges)
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
		next_perm(nodes)
		maxPerms -= 1
	}

	fmt.Println("Part 1:", min)
	fmt.Println("Part 2:", max)
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	factorialOfNumber := number * factorial(number-1)

	return factorialOfNumber
}

func distance(d []int, edges edgeMap) int {
	total := 0
	for i := 0; i < len(d)-1; i++ {
		k := fmt.Sprintf("%d%d", d[i], d[i+1])
		total += edges[k]
	}

	return total
}

func next_perm(a []int) {
	i := len(a) - 1
	j := len(a)
	for a[i-1] >= a[i] {
		i -= 1
	}

	for a[j-1] <= a[i-1] {
		j -= 1
	}

	a[i-1], a[j-1] = a[j-1], a[i-1]
	i += 1
	j = len(a)

	for i < j {
		a[i-1], a[j-1] = a[j-1], a[i-1]
		i += 1
		j -= 1
	}
}

func parse(file string) (edgeMap, []int) {
	nodes := []int{}
	m := make(map[string]int)
	edgeMap := make(edgeMap)
	id := 0
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		re := regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)
		matches := re.FindAllSubmatch(line, -1)
		from := string(matches[0][1])
		to := string(matches[0][2])
		weight, _ := strconv.Atoi(string(matches[0][3]))
		_, fok := m[from]
		_, tok := m[to]

		if !fok {
			m[from] = id
			nodes = append(nodes, id)
			id++
		}

		if !tok {
			m[to] = id
			nodes = append(nodes, id)
			id++
		}

		lr := fmt.Sprintf("%d%d", m[from], m[to])
		rl := fmt.Sprintf("%d%d", m[to], m[from])
		edgeMap[lr] = weight
		edgeMap[rl] = weight
	}

	return edgeMap, nodes
}

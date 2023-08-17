package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type edgeMap = map[string]int
type route struct {
	from   string
	to     string
	weight int
}

func main() {
	routes := parse("input")
	edges, nodes := extractGraphEls(routes)
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

func extractGraphEls(routes []route) (edgeMap, []int) {
	nodes := []int{}
	m := make(map[string]int)
	edgeMap := make(edgeMap)
	id := 0

	for _, r := range routes {
		s := []string{r.to, r.from}

		for _, sp := range s {
			_, ok := m[sp]
			if !ok {
				m[sp] = id
				nodes = append(nodes, id)
				id++
			}
		}
	}

	for _, r := range routes {
		lr := fmt.Sprintf("%d%d", m[r.from], m[r.to])
		rl := fmt.Sprintf("%d%d", m[r.to], m[r.from])
		edgeMap[lr] = r.weight
		edgeMap[rl] = r.weight
	}

	return edgeMap, nodes
}

func parse(file string) []route {
	routes := []route{}
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		sp := strings.Split(line, " ")

		from, to := sp[0], sp[2]
		weight, _ := strconv.Atoi(sp[4])
		routes = append(routes, route{from, to, weight})
	}

	return routes
}

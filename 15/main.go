package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	ingredients := parse("input")
	fmt.Println("Part 1:", balancing(ingredients))
}

func balancing(ingredients []ingredient) int {
	max := 0
	spoons := 100
	l := len(ingredients)
	groups := fixedLengthPartitions(spoons, l)

	for _, g := range groups {
		nMax := calculate(g, ingredients)
		if nMax > max {
			max = nMax
		}
	}
	return max
}

func calculate(a []int, ingredients []ingredient) int {
	totalI := ingredient{name: "Total"}

	for i, amount := range a {
		ii := ingredients[i]
		totalI.capacity += ii.capacity * amount
		totalI.durability += ii.durability * amount
		totalI.flavor += ii.flavor * amount
		totalI.texture += ii.texture * amount
	}

	return max(totalI.capacity, 0) *
		max(totalI.durability, 0) *
		max(totalI.flavor, 0) *
		max(totalI.texture, 0)
}

func fixedLengthPartitions(n, s int) [][]int {
	var partitions [][]int

	var generatePartitions func(remaining, length int, partition []int)
	generatePartitions = func(remaining, length int, partition []int) {
		if length == s {
			if remaining == 0 {
				partitions = append(partitions, append([]int(nil), partition...))
			}
			return
		}

		for i := 0; i <= min(remaining, n); i++ {
			partition[length] = i
			generatePartitions(remaining-i, length+1, partition)
		}
	}

	generatePartitions(n, 0, make([]int, s))
	return partitions
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parse(file string) []ingredient {
	ingredients := []ingredient{}
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
			`([A-Za-z]+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`,
		)
		matches := re.FindAllSubmatch(line, -1)
		name := string(matches[0][1])

		capacity, _ := strconv.Atoi(string(matches[0][2]))
		durability, _ := strconv.Atoi(string(matches[0][3]))
		flavor, _ := strconv.Atoi(string(matches[0][4]))
		texture, _ := strconv.Atoi(string(matches[0][5]))
		calories, _ := strconv.Atoi(string(matches[0][6]))

		i := ingredient{name, capacity, durability, flavor, texture, calories}
		ingredients = append(ingredients, i)
	}

	return ingredients
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type calcHeapFunc = func(a []int, i []ingredient) int

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
	groups := fixedPartition(spoons, l)

	for _, g := range groups {
		heapPermutation(g, l, ingredients, permFunc, &max)
	}
	return max
}

func permFunc(a []int, ingredients []ingredient) int {
	totalI := ingredient{name: "Total"}

	for i, amount := range a {
		ii := ingredients[i]
		totalI.capacity += ii.capacity * amount
		totalI.durability += ii.durability * amount
		totalI.flavor += ii.flavor * amount
		totalI.texture += ii.texture * amount
	}

	// This should be easier ...
	if totalI.capacity < 0 {
		totalI.capacity = 0
	}
	if totalI.durability < 0 {
		totalI.durability = 0
	}
	if totalI.flavor < 0 {
		totalI.flavor = 0
	}
	if totalI.texture < 0 {
		totalI.texture = 0
	}

	return totalI.capacity * totalI.durability * totalI.flavor * totalI.texture
}

// Stolen from some Python code and translated it by hand
func fixedPartition(n int, l int) [][]int {
	groups := [][]int{}
	partition := []int{n - l + 1}
	for i := 0; i < l-1; i++ {
		partition = append(partition, 1)
	}
	for {
		copyGroup := make([]int, len(partition))
		copy(copyGroup, partition)
		groups = append(groups, copyGroup)

		if partition[0]-1 > partition[1] {
			partition[0] -= 1
			partition[1] += 1
			continue
		}
		j := 2
		s := partition[0] + partition[1] - 1
		for j < l && partition[j] >= partition[0]-1 {
			s += partition[j]
			j += 1
		}
		if j >= l {
			break
		}
		x := partition[j] + 1
		partition[j] = x
		j -= 1
		for j > 0 {
			partition[j] = x
			s -= x
			j -= 1
		}
		partition[0] = s
	}
	return groups
}

func heapPermutation(
	a []int,
	size int,
	ingredients []ingredient,
	calc calcHeapFunc,
	max *int,
) {
	if size == 1 {
		newMax := calc(a, ingredients)
		if newMax > *max {
			*max = newMax
		}
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1, ingredients, calc, max)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
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

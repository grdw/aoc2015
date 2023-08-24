package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	buckets := parse(content)
	result := powerSet(buckets)
	fmt.Println("Part 1:", pourLiters(result, 150))
	fmt.Println("Part 2:", pourMinimumLiters(result, 150))
}

func pourLiters(bucketSet [][]int, liters int) int {
	count := 0

	for _, subset := range bucketSet {
		if totalSubset(subset) == liters {
			count++
		}
	}

	return count
}

func pourMinimumLiters(bucketSet [][]int, liters int) int {
	count := 0
	minLen := math.MaxInt64
	for _, subset := range bucketSet {
		l := len(subset)
		if totalSubset(subset) == liters && l < minLen {
			minLen = l
		}
	}

	for _, subset := range bucketSet {
		if totalSubset(subset) == liters && len(subset) == minLen {
			count++
		}
	}

	return count
}

func totalSubset(subset []int) int {
	total := 0
	for _, s := range subset {
		total += s
	}
	return total
}

func powerSet(nums []int) [][]int {
	var result [][]int
	generatePowerSet(nums, 0, []int{}, &result)
	return result
}

func generatePowerSet(nums []int, index int, current []int, result *[][]int) {
	if index == len(nums) {
		*result = append(*result, append([]int(nil), current...))
		return
	}

	// Include the current element in the subset
	current = append(current, nums[index])
	generatePowerSet(nums, index+1, current, result)

	// Exclude the current element from the subset
	current = current[:len(current)-1]
	generatePowerSet(nums, index+1, current, result)
}

func parse(input []byte) []int {
	s := bytes.Split(input, []byte("\n"))
	arr := []int{}
	for _, p := range s {
		i, err := strconv.Atoi(string(p))
		if err != nil {
			continue
		}
		arr = append(arr, i)
	}
	return arr
}

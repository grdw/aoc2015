package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	l int
	w int
	h int
}

func main() {
	cubes, err := parse("input")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer to part 1 is: %d\n", wrappingPaper(cubes))
	fmt.Printf("The answer to part 2 is: %d\n", ribbonLength(cubes))
}

func wrappingPaper(cubes []cube) int {
	total := 0

	for _, c := range cubes {
		lw := c.l * c.w
		wh := c.w * c.h
		hl := c.h * c.l
		slack := min([]int{lw, wh, hl})
		total += (lw * 2) + (wh * 2) + (hl * 2) + slack
	}

	return total
}

func ribbonLength(cubes []cube) int {
	length := 0

	for _, c := range cubes {
		l := []int{c.l, c.w, c.h}
		l[maxIn(l)] = 0

		for _, ll := range l {
			length += ll * 2
		}

		length += c.l * c.w * c.h
	}
	return length
}

func parse(file string) ([]cube, error) {
	cubes := []cube{}
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		return cubes, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		t := strings.Split(line, "x")
		l, _ := strconv.Atoi(t[0])
		w, _ := strconv.Atoi(t[1])
		h, _ := strconv.Atoi(t[2])

		cubes = append(cubes, cube{l, w, h})
	}

	return cubes, nil
}

func min(x []int) int {
	m := math.MaxInt32
	for _, n := range x {
		if n < m {
			m = n
		}
	}
	return m
}

func maxIn(x []int) int {
	m := 0
	j := 0
	for i, n := range x {
		if n > m {
			m = n
			j = i
		}
	}
	return j
}

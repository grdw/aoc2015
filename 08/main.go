package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1:", parse("input"))
}

func parse(file string) int {
	left := 0
	right := 0
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		bs, _ := strconv.Unquote(line)
		left += len(line)
		right += len(bs)
	}

	return left - right
}

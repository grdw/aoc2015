package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const s int = 1000
const turnOn string = "turn on"
const turnOff string = "turn off"
const toggle string = "toggle"

type grid = [s][s]int
type gridTransformFunc = func(grid *grid, instruction string, y int, x int)

func main() {
	g1 := parse("input", applyLights)
	fmt.Println("Part 1:", countLitLights(&g1))

	g2 := parse("input", applyBrightness)
	fmt.Println("Part 2:", countBrightness(&g2))
}

func applyBrightness(grid *grid, instruction string, y int, x int) {
	switch instruction {
	case turnOn:
		grid[y][x] += 1
		break
	case turnOff:
		if grid[y][x] > 0 {
			grid[y][x] -= 1
		}
		break
	case toggle:
		grid[y][x] += 2
		break
	}
}

func countBrightness(grid *grid) int {
	b := 0
	for y := 0; y < s; y++ {
		for x := 0; x < s; x++ {
			b += grid[y][x]
		}
	}
	return b
}

func applyLights(grid *grid, instruction string, y int, x int) {
	switch instruction {
	case turnOn:
		grid[y][x] = 1
		break
	case turnOff:
		grid[y][x] = 0
		break
	case toggle:
		if grid[y][x] == 0 {
			grid[y][x] = 1
		} else {
			grid[y][x] = 0
		}
		break
	}
}

func countLitLights(grid *grid) int {
	count := 0
	for y := 0; y < s; y++ {
		for x := 0; x < s; x++ {
			if grid[y][x] == 1 {
				count++
			}
		}
	}
	return count
}

func parse(file string, gridTransformer gridTransformFunc) grid {
	var g grid
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
			`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`,
		)

		matches := re.FindAllSubmatch(line, -1)
		t := string(matches[0][1])
		sx, _ := strconv.Atoi(string(matches[0][2]))
		sy, _ := strconv.Atoi(string(matches[0][3]))
		ex, _ := strconv.Atoi(string(matches[0][4]))
		ey, _ := strconv.Atoi(string(matches[0][5]))

		for y := sy; y <= ey; y++ {
			for x := sx; x <= ex; x++ {
				gridTransformer(&g, t, y, x)
			}
		}
	}

	return g
}

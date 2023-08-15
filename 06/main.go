package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const s int = 1000

type grid = [s][s]int
type instFn = func(grid *grid, y int, x int)

func main() {
	g1 := parse("input", turnOnLight, turnOffLight, toggleLight)
	fmt.Println("Part 1:", countLitLights(&g1))

	g2 := parse("input", addBrightness, reduceBrightness, addBrightnessByTwo)
	fmt.Println("Part 2:", countBrightness(&g2))
}

func addBrightness(grid *grid, y int, x int) {
	grid[y][x] += 1
}

func reduceBrightness(grid *grid, y int, x int) {
	if grid[y][x] > 0 {
		return
	}

	grid[y][x] -= 1
}

func addBrightnessByTwo(grid *grid, y int, x int) {
	grid[y][x] += 2
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

func turnOnLight(grid *grid, y int, x int) {
	grid[y][x] = 1
}

func turnOffLight(grid *grid, y int, x int) {
	grid[y][x] = 0
}

func toggleLight(grid *grid, y int, x int) {
	grid[y][x] ^= 1
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

func parse(file string, turnOn instFn, turnOff instFn, toggle instFn) grid {
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
				switch t {
				case "turn on":
					turnOn(&g, y, x)
					break
				case "turn off":
					turnOff(&g, y, x)
					break
				case "toggle":
					toggle(&g, y, x)
					break
				}
			}
		}
	}

	return g
}

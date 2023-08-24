package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x int
	y int
}

func main() {
	exGrid := parse("input_test", 6)
	fmt.Println("Ex 1:", animate(exGrid, 4, false))
	exGrid = parse("input_test", 6)
	fmt.Println("Ex 2:", animate(exGrid, 5, true))
	grid := parse("input", 100)
	fmt.Println("Part 1:", animate(grid, 100, false))
	grid = parse("input", 100)
	fmt.Println("Part 2:", animate(grid, 100, true))
}

func animate(grid [][]int, cycles int, locked bool) int {
	gridSize := len(grid)
	neighborsGrid := make([][]int, gridSize)

	if locked {
		grid[0][0] = 1
		grid[0][gridSize-1] = 1
		grid[gridSize-1][0] = 1
		grid[gridSize-1][gridSize-1] = 1
	}

	for i := 0; i < cycles; i++ {
		for y := 0; y < gridSize; y++ {
			neighborsGrid[y] = make([]int, gridSize)
			for x := 0; x < gridSize; x++ {
				neighbors := 0
				coords := []coord{
					coord{-1, -1},
					coord{-1, 1},
					coord{1, -1},
					coord{1, 1},
					coord{0, 1},
					coord{0, -1},
					coord{-1, 0},
					coord{1, 0},
				}

				for _, c := range coords {
					dx := x + c.x
					dy := y + c.y

					if dx < 0 || dx >= gridSize ||
						dy < 0 || dy >= gridSize {
						continue
					}

					if grid[dy][dx] == 1 {
						neighbors++
					}
				}

				neighborsGrid[y][x] = neighbors
			}
		}

		for y := 0; y < gridSize; y++ {
			for x := 0; x < gridSize; x++ {
				if locked && ((y == 0 && x == 0) ||
					(y == 0 && x == gridSize-1) ||
					(y == gridSize-1 && x == 0) ||
					(y == gridSize-1 && x == gridSize-1)) {
					continue
				}
				if grid[y][x] == 1 &&
					!(neighborsGrid[y][x] == 2 ||
						neighborsGrid[y][x] == 3) {
					grid[y][x] = 0
				} else if neighborsGrid[y][x] == 3 {
					grid[y][x] = 1
				}
			}
		}
	}
	return lightsOn(grid)
}

func lightsOn(grid [][]int) int {
	gridSize := len(grid)
	on := 0
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[y][x] == 1 {
				on++
			}
		}
	}

	return on
}

func parse(input string, gridSize int) [][]int {
	g := make([][]int, gridSize)

	readFile, err := os.Open(input)
	defer readFile.Close()

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	y := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		gridLine := make([]int, gridSize)
		for i, dot := range line {
			t := 0
			if dot == '#' {
				t = 1
			}
			gridLine[i] = t
		}
		g[y] = gridLine
		y++
	}

	return g
}

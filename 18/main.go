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
	fmt.Println("Ex 1:", animate(exGrid, 4))
	grid := parse("input", 100)
	fmt.Println("Part 1:", animate(grid, 100))
}

func animate(grid [][]int, cycles int) int {
	gridSize := len(grid)
	neighborsGrid := make([][]int, gridSize)

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

func printGrid(grid [][]int) {
	gridSize := len(grid)
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func printLightGrid(grid [][]int) {
	gridSize := len(grid)
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
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

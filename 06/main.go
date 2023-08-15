package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const s int = 1000

type instruction struct {
	t      string
	startX int
	startY int
	endX   int
	endY   int
}

func main() {
	instructions, err := parse("input")
	if err != nil {
		log.Fatal(err)
	}

	var grid [s][s]int
	applyLights(&grid, instructions)
	fmt.Println("Part 1:", countLitLights(&grid))
}

func applyLights(grid *[s][s]int, instructions []instruction) {
	for _, in := range instructions {
		for y := in.startY; y <= in.endY; y++ {
			for x := in.startX; x <= in.endX; x++ {
				switch in.t {
				case "turn on":
					grid[y][x] = 1
					break
				case "turn off":
					grid[y][x] = 0
					break
				case "toggle":
					if grid[y][x] == 0 {
						grid[y][x] = 1
					} else {
						grid[y][x] = 0
					}
					break
				}
			}
		}
	}
}

func countLitLights(grid *[s][s]int) int {
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

func parse(file string) ([]instruction, error) {
	instructions := []instruction{}
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		return instructions, err
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

		instructions = append(instructions, instruction{
			t, sx, sy, ex, ey,
		})
	}

	return instructions, nil
}

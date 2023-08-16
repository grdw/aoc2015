package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type umap = map[string]uint16
type operation struct {
	id       string
	operand  string
	executed bool
}

func main() {
	operations, values := parse("input")
	longSolve(operations, values)
	fmt.Println("Solution to Part 1:", values["a"])
}

func longSolve(operations []operation, values umap) {
	ex := len(values) + len(operations)

	for len(values) < ex {
		solveWires(operations, values)
	}
}

func solveWires(operations []operation, values umap) {
	for _, o := range operations {
		if o.executed {
			continue
		}

		s := strings.Split(o.operand, " ")
		if isLower(s[0]) && len(s) > 1 {
			l, lok := values[s[0]]
			r, rok := values[s[2]]
			lv, lerr := strconv.Atoi(s[0])
			rv, rerr := strconv.Atoi(s[2])

			if lerr == nil {
				l = uint16(lv)
			}

			if rerr == nil {
				r = uint16(rv)
			}

			if !((lok || lerr == nil) && (rok || rerr == nil)) {
				continue
			}

			switch s[1] {
			case "AND":
				values[o.id] = l & r
				break
			case "OR":
				values[o.id] = l | r
				break
			case "LSHIFT":
				values[o.id] = l << rv
				break
			case "RSHIFT":
				values[o.id] = l >> rv
				break
			}
			o.executed = true
		} else if s[0] == "NOT" {
			val, ok := values[s[1]]

			if ok {
				values[o.id] = ^val
			}
		} else {
			val, ok := values[s[0]]

			if ok {
				values[o.id] = val
			}
		}
	}
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func parse(file string) ([]operation, umap) {
	operations := []operation{}
	values := make(umap)
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		matches := strings.Split(line, " -> ")

		value, err := strconv.Atoi(matches[0])
		id := matches[1]
		if err != nil {
			operations = append(operations, operation{id, matches[0], false})
		} else {
			values[id] = uint16(value)
		}
	}

	return operations, values
}

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
	executed bool
	left     string
	right    string
	op       string
}

func main() {
	operations, values := parse("input")
	longSolve(operations, values)
	solution := values["a"]
	fmt.Println("Solution to part 1:", solution)

	operations, values = parse("input")
	values["b"] = solution
	longSolve(operations, values)
	fmt.Println("Solution to part 2:", values["a"])
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

		l, lok := values[o.left]
		r, rok := values[o.right]
		lv, lerr := strconv.Atoi(o.left)
		rv, rerr := strconv.Atoi(o.right)

		if lerr == nil {
			l = uint16(lv)
		}

		if rerr == nil {
			r = uint16(rv)
		}

		// Tests if the gate is closed
		if !((lok || lerr == nil || o.op == "NOT" || o.op == "ASSIGN") && (rok || rerr == nil)) {
			continue
		}

		switch o.op {
		case "AND":
			values[o.id] = l & r
			break
		case "OR":
			values[o.id] = l | r
			break
		case "LSHIFT":
			values[o.id] = l << r
			break
		case "RSHIFT":
			values[o.id] = l >> r
			break
		case "NOT":
			values[o.id] = ^r
			break
		case "ASSIGN":
			values[o.id] = r
			break
		}
		o.executed = true
	}
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
			left, right, op := parseOp(matches[0])

			operations = append(
				operations,
				operation{id, false, left, right, op},
			)
		} else {
			values[id] = uint16(value)
		}
	}

	return operations, values
}

func parseOp(op string) (string, string, string) {
	left := ""
	s := strings.Split(op, " ")

	if isLower(s[0]) && len(s) > 1 {
		return s[0], s[2], s[1]
	} else if s[0] == "NOT" {
		return left, s[1], "NOT"
	} else {
		return left, s[0], "ASSIGN"
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

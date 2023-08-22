package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", parseInput(content))
	fmt.Println("Part 2:", parseInputNoRed(content))
}

func parseInput(text []byte) int {
	total := 0
	blurb := []byte{}
	for _, b := range text {
		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			blurb = append(blurb, b)
			break
		default:
			i, _ := strconv.Atoi(string(blurb))
			total += i
			blurb = []byte{}
			break
		}
	}
	return total
}

func parseInputNoRed(text []byte) int {
	var values interface{}
	err := json.Unmarshal(text, &values)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	recurseParse(values, &total)
	return total
}

func recurseParse(value interface{}, total *int) {
	switch v := value.(type) {
	case float64:
		*total += int(v)
		break
	case map[string]interface{}:
		isRed := false
		for _, red := range v {
			if red == "red" {
				isRed = true
				break
			}
		}
		if !isRed {
			for _, vv := range v {
				recurseParse(vv, total)
			}
		}
		break
	case []interface{}:
		for _, vv := range v {
			recurseParse(vv, total)
		}
		break
	}
}

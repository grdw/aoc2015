package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

type replacements = map[string][]string

func main() {
	reps, input := parse("input")
	possibilities := molecules(input, reps)
	fmt.Println("Part 1:", len(possibilities))
	fmt.Println("Part 2:", recGenMolecule(input, reps))
}

func recGenMolecule(input string, reps replacements) int {
	res := math.MaxInt32
	ml := math.MaxInt32
	flippedReps := make(map[string]string)
	for k, v := range reps {
		for _, vv := range v {
			flippedReps[vv] = k
		}
	}

	genMolecule(input, "e", flippedReps, 0, &res, &ml)
	return res
}

func genMolecule(
	start string,
	input string,
	reps map[string]string,
	cycle int,
	res *int,
	minLen *int) {

	if input == start && cycle < *res {
		*res = cycle
	}

	list := make(map[string]bool)
	for k, rep := range reps {
		re := regexp.MustCompile(k)
		borders := re.FindAllStringIndex(start, -1)

		for _, border := range borders {
			result := fmt.Sprintf(
				"%s%s%s",
				start[:border[0]],
				rep,
				start[border[1]:],
			)
			list[result] = true
			q := len(result)
			if q < *minLen {
				*minLen = q
			}
		}
	}

	for l, _ := range list {
		if len(l) == *minLen {
			genMolecule(l, input, reps, cycle+1, res, minLen)
		}
	}
}

func molecules(input string, reps replacements) []string {
	unique := make(map[string]bool)
	for k, rep := range reps {
		re := regexp.MustCompile(k)
		borders := re.FindAllStringIndex(input, -1)

		for _, r := range rep {
			for _, border := range borders {
				result := fmt.Sprintf(
					"%s%s%s",
					input[:border[0]],
					r,
					input[border[1]:],
				)
				unique[result] = true
			}
		}
	}

	m := []string{}
	for k, _ := range unique {
		m = append(m, k)
	}
	return m
}

func parse(input string) (replacements, string) {
	reps := make(replacements)
	i := ""
	readFile, err := os.Open(input)
	defer readFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.Contains(line, " => ") {
			rawR := strings.Split(line, " => ")
			reps[rawR[0]] = append(reps[rawR[0]], rawR[1])
		} else if line != "" {
			i = line
		}
	}

	return reps, i
}

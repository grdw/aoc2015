package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type replacements = map[string][]string

func main() {
	reps, input := parse("input")
	possibilities := molecules(input, reps)
	fmt.Println("Part 1:", len(possibilities))
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

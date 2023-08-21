package main

import "fmt"

const input string = "vzbxkghb"
const min byte = 97
const max byte = min + 25

func increment(input []byte) []byte {
	output := input
	index := len(input) - 1
	for {
		if index < 0 {
			output = append([]byte{'a'}, output...)
			break
		} else if output[index]+1 < max {
			output[index]++
			break
		} else {
			output[index] = 97
			index--
		}
	}
	return output
}

func valid(input []byte) bool {
	return validTriple(input) && validLetters(input) && validDouble(input)
}

func validTriple(input []byte) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+1]-1 && input[i+1]-1 == input[i+2]-2 {
			return true
		}
	}
	return false
}

func validLetters(input []byte) bool {
	for _, l := range input {
		if l == 'i' || l == 'o' || l == 'l' {
			return false
		}
	}
	return true
}

func validDouble(input []byte) bool {
	count := 0
	prevInput := byte(0)
	ticked := make(map[byte]bool)

	for i := 0; i < len(input); i++ {
		_, ok := ticked[input[i]]
		if input[i] == prevInput && !ok {
			count++
			ticked[input[i]] = true
		}

		prevInput = input[i]
	}

	return count == 2
}

func findNextPassword(input []byte) []byte {
	for {
		input := increment(input)
		if valid(input) {
			return input
		}
	}
}

func main() {
	fmt.Println("vim-go")
}

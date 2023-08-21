package main

import "fmt"

const input string = "vzbxkghb"
const min byte = 97
const max byte = min + 26

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
			output[index] = 'a'
			index--
		}
	}
	return output
}

func valid(input []byte) bool {
	return validTriple(input) && validDoubleAndLetters(input)
}

func validTriple(input []byte) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+1]-1 && input[i+1]-1 == input[i+2]-2 {
			return true
		}
	}
	return false
}

func validDoubleAndLetters(input []byte) bool {
	count := 0
	prevInput := byte(0)
	doublePairChar := byte(0)

	for _, l := range input {
		if l == 'i' || l == 'o' || l == 'l' {
			return false
		}

		if l == prevInput && doublePairChar != l {
			count++
			doublePairChar = l
		}

		prevInput = l
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
	fmt.Println("part 1:", string(findNextPassword([]byte(input))))
	fmt.Println("part 2:", string(findNextPassword(findNextPassword([]byte(input)))))
}

package main

import "testing"

func TestHouses(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{input: ">", answer: 0},
		{input: "^>v<", answer: 1},
		{input: "^v^v^v^v^v", answer: 2},
	}

	for _, test := range tests {
		n := houses([]byte(test.input))
		if n != test.answer {
			t.Fatal("Incorrect")
		}
	}
}

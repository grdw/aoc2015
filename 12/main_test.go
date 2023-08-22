package main

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	type test struct {
		input  string
		output int
	}

	tests := []test{
		{
			input:  "[1,2,3]",
			output: 6,
		},
		{
			input:  "{\"a\":2,\"b\":4}",
			output: 6,
		},
		{
			input:  "[[[3]]]",
			output: 3,
		},
		{
			input:  "{\"a\":{\"b\":4},\"c\":-1}",
			output: 3,
		},
		{
			input:  "{\"a\":[-1,1]}",
			output: 0,
		},
	}

	for _, tt := range tests {
		out := parseInput([]byte(tt.input))
		if out != tt.output {
			t.Fatalf("Expected %d got %d", tt.output, out)
		}
	}
}

func TestParseInputNoRed(t *testing.T) {
	type test struct {
		input  string
		output int
	}

	tests := []test{
		{
			input:  "[1,2,3]",
			output: 6,
		},
		{
			input:  "[1,{\"c\":\"red\",\"b\":2},3]",
			output: 4,
		},
		{
			input:  "[1,{\"c\":\"blue\",\"b\":2},3]",
			output: 6,
		},
		{
			input:  "[{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}]",
			output: 0,
		},
		{
			input:  "[1,\"red\",5]",
			output: 6,
		},
	}

	for _, tt := range tests {
		out := parseInputNoRed([]byte(tt.input))
		if out != tt.output {
			t.Fatalf("Expected %d got %d", tt.output, out)
		}
	}
}

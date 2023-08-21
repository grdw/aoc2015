package main

import (
	"testing"
)

func TestLookAndSaySequence(t *testing.T) {
	type test struct {
		input  string
		output string
	}

	tests := []test{
		{input: "1", output: "11"},
		{input: "11", output: "21"},
		{input: "21", output: "1211"},
		{input: "1211", output: "111221"},
		{input: "111221", output: "312211"},
	}

	for _, tt := range tests {
		r := rbuv{[]byte(tt.input), 0}
		w := wbuv{[]byte{}}
		lookAndSaySequence(&r, &w)
		output := string(w.data)

		if output != tt.output {
			t.Fatalf("Got %s expected %s", output, tt.output)
		}
	}
}

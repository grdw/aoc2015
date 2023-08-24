package main

import "testing"

func TestMolecules(t *testing.T) {
	m := make(replacements)
	m["H"] = []string{"HO", "OH"}
	m["O"] = []string{"HH"}

	type test struct {
		input  string
		output int
	}

	tests := []test{
		test{"HOH", 4},
		test{"HOHOHO", 7},
	}

	for _, tt := range tests {
		mcules := molecules(tt.input, m)
		if len(mcules) != tt.output {
			t.Fatal("Nop")
		}
	}
}

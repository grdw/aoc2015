package main

import (
	"testing"
)

func TestMolecules(t *testing.T) {
	m := make(replacements)
	m["HO"] = "H"
	m["OH"] = "H"
	m["HH"] = "O"

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

func TestGenMolecule(t *testing.T) {
	m := make(replacements)
	m["H"] = "e"
	m["O"] = "e"
	m["HO"] = "H"
	m["OH"] = "H"
	m["HH"] = "O"

	type test struct {
		input  string
		output int
	}

	tests := []test{
		test{"HOH", 3},
		test{"HOHOHO", 6},
	}

	for _, tt := range tests {
		max := recGenMolecule(tt.input, m)
		if max != tt.output {
			t.Fatal("Nop")
		}
	}
}

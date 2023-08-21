package main

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	type test struct {
		input  []byte
		output string
	}
	tests := []test{
		{
			input:  []byte("a"),
			output: "b",
		},
		{
			input:  []byte("abz"),
			output: "aca",
		},
		{
			input:  []byte("zz"),
			output: "aaa",
		},
		{
			input:  []byte("az"),
			output: "ba",
		},
		{
			input:  []byte("vzbxxyyy"),
			output: "vzbxxyyz",
		},
	}

	for _, test := range tests {
		out := string(increment(test.input))
		if out != test.output {
			t.Fatalf("Expected %s got %s\n", test.output, out)
		}
	}
}

func TestValidate(t *testing.T) {
	type test struct {
		input []byte
		valid bool
	}

	tests := []test{
		{
			input: []byte("hijklmmn"),
			valid: false,
		},
		{
			input: []byte("abbceffg"),
			valid: false,
		},
		{
			input: []byte("abccegjk"),
			valid: false,
		},
		{
			input: []byte("abcccegjk"),
			valid: false,
		},
		{
			input: []byte("abcdffaa"),
			valid: true,
		},
		{
			input: []byte("vzbxxyzz"),
			valid: true,
		},
	}

	for _, test := range tests {
		o := valid(test.input)
		if o != test.valid {
			t.Fatalf("Expected %t but got %t for %s", test.valid, o, test.input)
		}
	}
}

func TestNextPassword(t *testing.T) {
	type test struct {
		input  []byte
		output string
	}

	tests := []test{
		{
			input:  []byte("abcdefgh"),
			output: "abcdffaa",
		},
		{
			input:  []byte("ghijklmn"),
			output: "ghjaabcc",
		},
	}

	for _, test := range tests {
		out := findNextPassword(test.input)
		if string(out) != test.output {
			t.Fatalf("Expected %s but got %s", string(test.input), test.output)
		}
	}
}

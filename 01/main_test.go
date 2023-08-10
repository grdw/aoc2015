package main

import (
	"testing"
)

func TestCountFloors(t *testing.T) {
	tests := []struct {
		floors  int
		content string
	}{
		{floors: 0, content: "()()"},
		{floors: 1, content: "("},
		{floors: 1, content: "(()"},
	}

	for _, tc := range tests {
		n := countFloors([]byte(tc.content))
		if n != tc.floors {
			t.Fatal("Thing didn't match")
		}
	}
}

func TestBasement(t *testing.T) {
	tests := []struct {
		basementIndex int
		content       string
	}{
		{basementIndex: 1, content: ")"},
		{basementIndex: 5, content: "()())"},
	}

	for _, tc := range tests {
		n := basement([]byte(tc.content))
		if n != tc.basementIndex {
			t.Fatal("Thing didn't match")
		}
	}
}

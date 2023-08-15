package main

import (
	"testing"
)

func TestPartOneEx1(t *testing.T) {
	result := calculateMd5("abcdef", "00000")
	if result != 609043 {
		t.Fatal("Does not match")
	}
}

func TestPartOneEx2(t *testing.T) {
	result := calculateMd5("pqrstuv", "00000")
	if result != 1048970 {
		t.Fatal("Does not match")
	}
}

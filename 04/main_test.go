package main

import (
	"fmt"
	"testing"
)

func TestPartOneEx1(t *testing.T) {
	result := p1("abcdef")
	fmt.Println(result)
	if result != 609043 {
		t.Fatal("Does not match")
	}
}

func TestPartOneEx2(t *testing.T) {
	result := p1("pqrstuv")
	fmt.Println(result)
	if result != 1048970 {
		t.Fatal("Does not match")
	}
}

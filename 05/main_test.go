package main

import "testing"

func TestTwins(t *testing.T) {
	l := twins("qjhvhtzxzqqjkmpb")
	if !l {
		t.Fatal("There are twins")
	}

	l = twins("xxyxx")
	if !l {
		t.Fatal("There are twins")
	}
}

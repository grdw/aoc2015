package main

import "testing"

func TestRibbonLength(t *testing.T) {
	n := ribbonLength([]byte("1x1x10"))
	if n != 14 {
		t.Fatal("boom")
	}
}

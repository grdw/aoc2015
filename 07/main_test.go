package main

import (
	"testing"
)

func TestTreeParsing(t *testing.T) {
	ops, values := parse("input_test")
	if len(ops) != 6 {
		t.Fatal("Nodes length is too short")
	}

	if ops[0].id != "d" {
		t.Fatal("Node id is incorrect")
	}

	if values["x"] != 123 {
		t.Fatal("Node value is incorrect")
	}

	if ops[0].operand != "x AND y" {
		t.Fatal("Node operand is incorrect")
	}

	longSolve(ops, values)
	answers := make(map[string]uint16)
	answers["d"] = 72
	answers["e"] = 507
	answers["f"] = 492
	answers["g"] = 114
	answers["h"] = 65412
	answers["i"] = 65079
	answers["x"] = 123
	answers["y"] = 456

	for k, a := range answers {
		if a != values[k] {
			t.Fatalf("%s %d does not match %d\n", k, a, values[k])
		}
	}
}

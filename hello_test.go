package main

import "testing"

func TestAddTwoInts(t *testing.T) {
	a := addTwoInts(1, 3)
	if a != 4 {
		t.Error("1 + 3 didn't equal 4")
	}
}

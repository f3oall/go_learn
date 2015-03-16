package main

import "testing"

func TestScan(t *testing.T) {
	max := 5
	choice := Scan(max)
	if choice > 5 {
		t.Errorf("Value is bigger than maximum")
	}
	if choice <= 0 {
		t.Errorf("Value is equal to 0, or less than zero")
	}
}

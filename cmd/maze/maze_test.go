package main

import "testing"

func TestSum(t *testing.T) {
	r := sum(3, 4)
	if r != 7 {
		t.Error("should be seven.")
	}
}

package builders

import (
	"testing"
)

func TestRandomVal(t *testing.T) {
	reInitRand(0)
	if getRand(5) != 4 {
		t.Error("getRand(5) not returning as expected, should be 4")
	}
	if getRand(5) != 4 {
		t.Error("getRand(5) not returning as expected, should be 4")
	}
	if getRand(5) != 3 {
		t.Error("getRand(5) not returning as expected, should be 3")
	}
}

func TestRandomZero(t *testing.T) {
	if getRand(0) != 0 {
		t.Error("Random of zero should return zero")
	}
}

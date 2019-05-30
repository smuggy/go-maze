package basics

import (
	"testing"
)

func TestRandomVal(t *testing.T) {
	ReInitRand(0)
	if GetRand(5) != 4 {
		t.Error("getRand(5) not returning as expected, should be 4")
	}
	if GetRand(5) != 4 {
		t.Error("getRand(5) not returning as expected, should be 4")
	}
	if GetRand(5) != 3 {
		t.Error("getRand(5) not returning as expected, should be 3")
	}
}

func TestRandomValInt8(t *testing.T) {
	ReInitRand(0)
	if GetRand8(5) != 4 {
		t.Error("getRand(5) not returning as expected, should be 4")
	}
	if GetRand8(5) != 4 {
		t.Error("getRand(5) not returning as expected, should be 4")
	}
	if GetRand8(5) != 3 {
		t.Error("getRand(5) not returning as expected, should be 3")
	}
}

func TestRandomZero(t *testing.T) {
	if GetRand(0) != 0 {
		t.Error("Random of zero should return zero")
	}
}

func TestRandomZeroInt8(t *testing.T) {
	if GetRand8(0) != 0 {
		t.Error("Random of zero should return zero")
	}
}

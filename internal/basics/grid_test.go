package basics

import (
	"testing"
)

func TestSize(t *testing.T) {
	g := BuildGrid(5, 5)
	if g.Size() != 25 {
		t.Error("Size should be 25...")
	}
}

func TestBuild(t *testing.T) {
	g := BuildGrid(2, 2)
	c := g.get(0, 0)
	if c.North != nil {
		t.Error("North should be nil...")
	}
	if c.East == nil {
		t.Error("East should not be nil")
	}
	if c.West != nil {
		t.Error("West should be nil...")
	}
	if c.South == nil {
		t.Error("South should not be nil...")
	}
	if c.South.equals(*g.get(1, 0)) == false {
		t.Error("South should be 1,0")
	}
	c = g.get(0, 1)
	if c.North != nil {
		t.Error("North should be nil...")
	}
	if c.East != nil {
		t.Error("East should be nil...")
	}
	if c.West == nil {
		t.Error("West should not be nil...")
	}
	if c.West.equals(*g.get(0,0,)) == false {
		t.Error("West should be 0, 0")
	}
	if c.South == nil {
		t.Error("South should not be nil...")
	}
	if c.South.equals(*g.get(1, 1)) == false {
		t.Error("South should be 1,1")
	}

}

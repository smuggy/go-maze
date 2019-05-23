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
	if c.West.equals(*g.get(0, 0)) == false {
		t.Error("West should be 0, 0")
	}
	if c.South == nil {
		t.Error("South should not be nil...")
	}
	if c.South.equals(*g.get(1, 1)) == false {
		t.Error("South should be 1,1")
	}
}

func TestGrid_BuildGridEmpty(t *testing.T) {
	g := BuildGrid(0, 5)
	if g != nil {
		t.Error("zero rows should have not created grid")
	}
	h := BuildGrid(5, 0)
	if h != nil {
		t.Error("zero columns should have not created grid")
	}
}

func TestGrid_EachCell(t *testing.T) {
	var row int8 = 3
	var col int8 = 2
	g := BuildGrid(row, col)
	var i int8 = 0
	var j int8 = 0
	for c := range g.EachCell() {
		if c.x != i {
			t.Errorf("x should be %d, but was %d", i, c.x)
		}
		if c.y != j {
			t.Errorf("y should be %d, but was %d", j, c.y)
		}
		j++
		if j >= col {
			j = 0
			i++
		}
	}
}

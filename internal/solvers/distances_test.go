package solvers

import (
	"github.com/smuggy/go-maze/internal/basics"
	"testing"
)

func TestBuilderWithNil(t *testing.T) {
	d := buildDistances(nil)
	if d != nil {
		t.Error("build should have failed.")
	}
}

func TestBuildDistances(t *testing.T) {
	g := buildSimpleGrid()
	d := buildDistances(g.GetCell(0, 0))
	if d == nil {
		t.Error("should have built distances map.")
	} else {
		if len(d.cells) != 1 {
			t.Error("should have one item in cell list")
		}
		if d.GetCellDistance(g.GetCell(0, 0)) != 0 {
			t.Error("expecting zero distance")
		}
		if d.GetCellDistance(g.GetCell(0, 1)) != NoDistanceSet {
			t.Error("no distance was set.")
		}
		d.SetCellDistance(g.GetCell(0, 1), 1)
		if d.GetCellDistance(g.GetCell(0, 1)) != 1 {
			t.Error("distnace should be one.")
		}
	}
}

func TestBuildPathRoot(t *testing.T) {
	g := buildSimpleGrid()
	d := buildDistances(g.GetCell(0, 0))
	d.SetCellDistance(g.GetCell(0, 1), 1)
	b := d.BuildPath(g.GetCell(0, 0))
	if b.cells[g.GetCell(0, 0)] != 0 {
		t.Error("0,0 Value incorrectly set")
	}
}

func TestBuildPath(t *testing.T) {
	g := buildSimpleGrid()
	d := buildDistances(g.GetCell(0, 0))
	d.SetCellDistance(g.GetCell(0, 1), 1)
	b := d.BuildPath(g.GetCell(0, 1))
	if b.cells[g.GetCell(0, 1)] != 1 {
		t.Error("0,1 Value incorrectly set")
	}
}

func TestBuildApplyAll(t *testing.T) {
	g := buildSimpleGrid()
	d := buildDistances(g.GetCell(0, 0))
	d.SetCellDistance(g.GetCell(0, 1), 1)
	d.ApplyToCellValue()
	if g.GetCell(0, 0).Value != "  0" {
		t.Error("0,0 Value incorrectly set")
	}
	if g.GetCell(0, 1).Value != "  1" {
		t.Error("0,1 Value incorrectly set")
	}
}

func TestBuildApplyRoot(t *testing.T) {
	g := buildSimpleGrid()
	d := buildDistances(g.GetCell(0, 0))
	d.SetCellDistance(g.GetCell(0, 1), NoDistanceSet)
	d.ApplyToCellValue()
	if g.GetCell(0, 0).Value != "  0" {
		t.Error("0,0 Value incorrectly set")
	}

	if g.GetCell(0, 1).Value != "   " {
		t.Error("0,1 Value incorrectly set")
	}
}

func buildSimpleGrid() *basics.Grid {
	g := basics.BuildGrid(1, 2)
	g.GetCell(0, 1).Links = append(g.GetCell(0, 1).Links, g.GetCell(0, 0))
	g.GetCell(0, 0).Links = append(g.GetCell(0, 0).Links, g.GetCell(0, 1))
	return g
}

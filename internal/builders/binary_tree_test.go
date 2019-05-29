package builders

import (
	"github.com/smuggy/go-maze/internal/basics"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	g := basics.BuildGrid(2, 2)
	basics.ReInitRand(1)
	bt := BinaryTree{}
	bt.Build(g)
	if len(g.GetCell(0, 0).Links) != 1 ||
		g.GetCell(0, 0).IsLinked(g.GetCell(0, 1)) == false {
		t.Error("0, 0 should only have one link to east")
	}
	if len(g.GetCell(0, 1).Links) != 2 {
		t.Error("0, 1 should have two links")
	}
}
func TestNoGrid(t *testing.T) {
	var g *basics.Grid = nil
	bt := BinaryTree{}
	bt.Build(g)
}

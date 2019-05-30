package builders

import (
	"github.com/smuggy/go-maze/internal/basics"
)

type Sidewinder struct{}

func (s Sidewinder) Build(g *basics.Grid) {
	if g != nil {
		buildSidewinderMaze(g)
	}
}

func buildSidewinderMaze(grid *basics.Grid) {
	var run = make([]*basics.Cell, 0, 10)
	for c := range grid.EachCell() {
		if c.West == nil {
			// Start of new row...
			run = run[:0]
		}
		run = append(run, c)
		if c.East == nil || ((c.North != nil) && (basics.GetRand(2) == 0)) {
			var member = run[basics.GetRand(len(run)-1)]
			if member.North != nil {
				member.Link(member.North, true)
			}
			run = run[:0]
		} else {
			c.Link(c.East, true)
		}
	}
}

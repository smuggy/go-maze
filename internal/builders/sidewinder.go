package builders

import (
	"github.com/smuggy/go-maze/internal/basics"
)

func Sidewinder(grid *basics.Grid) {
	var run = make([]*basics.Cell, 0, 10)
	for c := range grid.EachCell() {
		if c.West == nil {
			run = run[:0]
		}
		run = append(run, c)
		if c.East == nil || ((c.North != nil) && (getRand(2) == 0)) {
			var member = run[getRand(len(run)-1)]
			if member.North != nil {
				member.Link(member.North, true)
			}
			run = run[:0]
		} else {
			c.Link(c.East, true)
		}
	}
}

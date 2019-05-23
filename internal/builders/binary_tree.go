package builders

import (
	"github.com/smuggy/go-maze/internal/basics"
)

func BuildMaze(grid *basics.Grid) {
	binaryTree(grid)
}

func binaryTree(grid *basics.Grid) {

	for cell := range grid.EachCell() {
		neighbors := make([]*basics.Cell, 0, 2)
		if cell.North != nil {
			neighbors = append(neighbors, cell.North)
		}
		if cell.East != nil {
			neighbors = append(neighbors, cell.East)
		}
		if len(neighbors) > 0 {
			i := getRand(len(neighbors))
			cell.Link(neighbors[i], true)
		}
	}
}

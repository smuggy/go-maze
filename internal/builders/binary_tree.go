package builders

import (
	"github.com/smuggy/go-maze/internal/basics"
)

type MazeBuilder interface {
	Build(grid *basics.Grid)
}
type BinaryTree struct{}

func (b BinaryTree) Build(grid *basics.Grid) {
	if grid != nil {
		buildBinaryTreeMaze(grid)
	}
}

func buildBinaryTreeMaze(grid *basics.Grid) {

	neighbors := make([]*basics.Cell, 0, 4)
	for cell := range grid.EachCell() {
		neighbors = neighbors[:0]
		if cell.North != nil {
			neighbors = append(neighbors, cell.North)
		}
		if cell.East != nil {
			neighbors = append(neighbors, cell.East)
		}
		if len(neighbors) > 0 {
			i := basics.GetRand(len(neighbors))
			cell.Link(neighbors[i], true)
		}
	}
}

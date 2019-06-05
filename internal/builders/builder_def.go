package builders

import "github.com/smuggy/go-maze/internal/basics"

type MazeBuilder interface {
	Build(grid *basics.Grid)
}

type Builder int

const (
	AldousAlgo     Builder = 0
	BinaryTreeAlgo Builder = 1
	SidewinderAlgo Builder = 2
)

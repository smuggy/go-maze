package builders

import (
	"github.com/smuggy/go-maze/internal/basics"
)

type Aldous struct{}

func (a Aldous) Build(grid *basics.Grid) {
	if grid != nil {
		BuildAldousMaze(grid)
	}
}

func BuildAldousMaze(grid *basics.Grid) {
	var currCell = grid.GetRandomCell()

	for grid.IsFull() == false {
		nextCell := getNextCell(currCell)
		if len(nextCell.Links) == 0 {
			currCell.Link(nextCell, true)
		}
		currCell = nextCell
	}
}

func getNextCell(curr *basics.Cell) *basics.Cell {
	var ret *basics.Cell = nil
	for ret == nil {
		num := basics.GetRand(4)
		if num == 0 {
			ret = curr.North
		} else if num == 1 {
			ret = curr.East
		} else if num == 2 {
			ret = curr.South
		} else if num == 3 {
			ret = curr.West
		}
	}
	return ret
}

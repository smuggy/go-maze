package solvers

import (
	"github.com/smuggy/go-maze/internal/basics"
)

func DijkstraSolver(start *basics.Cell) *basics.Distances {
	var gridDistances = basics.BuildDistances(start)
	frontier := make([]*basics.Cell, 0)
	frontier = append(frontier, start)

	for len(frontier) > 0 {
		frontier = findNewFrontier(frontier, gridDistances)
	}
	return gridDistances
}

func findNewFrontier(frontier []*basics.Cell, gridDistances *basics.Distances) []*basics.Cell {
	newFrontier := make([]*basics.Cell, 0)
	for _, f := range frontier {
		for _, linked := range f.Links {
			if gridDistances.GetCellDistance(linked) != basics.NoDistanceSet {
				continue
			}
			gridDistances.SetCellDistance(linked, gridDistances.GetCellDistance(f)+1)
			newFrontier = append(newFrontier, linked)
		}
	}
	return newFrontier
}

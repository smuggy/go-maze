package solvers

import (
	"github.com/smuggy/go-maze/internal/basics"
)

type Solution interface {
	ApplyDistances()
}

type DijkstraSolution struct {
	dist *distances
}

func DijkstraSolver(start *basics.Cell) DijkstraSolution {
	var gridDistances = buildDistances(start)
	frontier := make([]*basics.Cell, 0)
	frontier = append(frontier, start)

	for len(frontier) > 0 {
		frontier = findNewFrontier(frontier, gridDistances)
	}
	return DijkstraSolution{dist: gridDistances}
}

func findNewFrontier(frontier []*basics.Cell, gridDistances *distances) []*basics.Cell {
	newFrontier := make([]*basics.Cell, 0)
	for _, f := range frontier {
		for _, linked := range f.Links {
			if gridDistances.getCellDistance(linked) != noDistanceSet {
				continue
			}
			gridDistances.setCellDistance(linked, gridDistances.getCellDistance(f)+1)
			newFrontier = append(newFrontier, linked)
		}
	}
	return newFrontier
}

func (sol DijkstraSolution) ApplyDistances() {
	sol.dist.applyToCellValue()
}

package solvers

import (
	"github.com/smuggy/go-maze/internal/basics"
)

type Solution interface {
	ApplyDistances()
	MaxPath() Solution
}

type DijkstraSolution struct {
	dist  *distances
	start *basics.Cell
}

func DijkstraSolver(start *basics.Cell) DijkstraSolution {
	var gridDistances = buildDistances(start)
	frontier := make([]*basics.Cell, 0)
	frontier = append(frontier, start)

	for len(frontier) > 0 {
		frontier = findNewFrontier(frontier, gridDistances)
	}
	return DijkstraSolution{dist: gridDistances, start: start}
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

func (sol DijkstraSolution) MaxPath() Solution {
	var goal = sol.dist.getMaxCell(sol.start)
	var bc = sol.dist.buildPath(goal)
	return DijkstraSolution{dist: bc, start: sol.start}
}

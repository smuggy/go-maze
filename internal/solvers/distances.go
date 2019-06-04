package solvers

import (
	"fmt"
	"github.com/smuggy/go-maze/internal/basics"
)

const noDistanceSet int = -1

type distances struct {
	root  *basics.Cell
	cells map[*basics.Cell]int
}

func buildDistances(root *basics.Cell) *distances {
	if root == nil {
		return nil
	}
	var d = new(distances)
	d.root = root
	d.cells = make(map[*basics.Cell]int)
	d.cells[root] = 0
	return d
}

func (d *distances) getCellDistance(c *basics.Cell) int {
	if distance, ok := d.cells[c]; ok {
		return distance
	}
	return noDistanceSet
}

func (d *distances) setCellDistance(c *basics.Cell, distance int) {
	d.cells[c] = distance
}

func (d *distances) buildPath(goal *basics.Cell) *distances {
	var current = goal
	var breadcrumbs = buildDistances(d.root)
	breadcrumbs.setCellDistance(current, d.getCellDistance(current))

	for current.Equals(*d.root) == false {
		for _, neighbor := range current.Links {
			if d.cells[neighbor] < d.cells[current] {
				breadcrumbs.cells[neighbor] = d.cells[neighbor]
				current = neighbor
				break
			}
		}
	}

	return breadcrumbs
}

func (d *distances) applyToCellValue() {
	for c, i := range d.cells {
		if i == noDistanceSet {
			c.Value = "   "
		} else {
			s := fmt.Sprintf("%3d", i)
			c.Value = s
		}
	}
}

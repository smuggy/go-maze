package solvers

import (
	"fmt"
	"github.com/smuggy/go-maze/internal/basics"
)

const NoDistanceSet int = -1

type Distances struct {
	root  *basics.Cell
	cells map[*basics.Cell]int
}

func buildDistances(root *basics.Cell) *Distances {
	if root == nil {
		return nil
	}
	var d = new(Distances)
	d.root = root
	d.cells = make(map[*basics.Cell]int)
	d.cells[root] = 0
	return d
}

func (d *Distances) GetCellDistance(c *basics.Cell) int {
	if distance, ok := d.cells[c]; ok {
		return distance
	}
	return NoDistanceSet
}

func (d *Distances) SetCellDistance(c *basics.Cell, distance int) {
	d.cells[c] = distance
}

func (d *Distances) BuildPath(goal *basics.Cell) *Distances {
	var current = goal
	var breadcrumbs = buildDistances(d.root)
	breadcrumbs.SetCellDistance(current, d.GetCellDistance(current))

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

func (d *Distances) ApplyToCellValue() {
	for c, i := range d.cells {
		if i == NoDistanceSet {
			c.Value = "   "
		} else {
			s := fmt.Sprintf("%3d", i)
			c.Value = s
		}
	}
}

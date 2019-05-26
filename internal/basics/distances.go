package basics

import "fmt"

const NoDistanceSet int = -1

type Distances struct {
	root  *Cell
	cells map[*Cell]int
}

func BuildDistances(root *Cell) *Distances {
	if root == nil {
		return nil
	}
	var d = new(Distances)
	d.root = root
	d.cells = make(map[*Cell]int)
	d.cells[root] = 0
	return d
}

func (d *Distances) GetCellDistance(c *Cell) int {
	if distance, ok := d.cells[c]; ok {
		return distance
	}
	return NoDistanceSet
}

func (d *Distances) SetCellDistance(c *Cell, distance int) {
	d.cells[c] = distance
}

func (d *Distances) BuildPath(goal *Cell) *Distances {
	var current = goal
	var breadcrumbs = BuildDistances(d.root)
	breadcrumbs.SetCellDistance(current, d.GetCellDistance(current))

	for current.equals(*d.root) == false {
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

		fmt.Println(c.ToString(), "\t", i)
		if i == NoDistanceSet {
			c.Value = "   "
		} else {
			s := fmt.Sprintf("%3d", i)
			c.Value = s
		}

	}
}

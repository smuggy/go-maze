package basics

import (
	"fmt"
)

type Cell struct {
	x     int8
	y     int8
	Value string
	North *Cell
	South *Cell
	East  *Cell
	West  *Cell
	Links []*Cell
}

func BuildCell(x int8, y int8) *Cell {
	if x < 0 {
		return nil
	}
	if y < 0 {
		return nil
	}
	return &(Cell{x, y, "   ", nil, nil, nil, nil, make([]*Cell, 0, 5)})
}

func (c *Cell) Equals(other Cell) bool {
	if c.x != other.x {
		return false
	}
	if c.y != other.y {
		return false
	}
	return true
}

func find(cells []*Cell, val *Cell) bool {
	if val == nil {
		return false
	}
	if cells == nil {
		return false
	}
	for i := range cells {
		if cells[i].Equals(*val) == true {
			return true
		}
	}
	return false
}

func remove(cells []*Cell, val *Cell) []*Cell {
	var newCells = make([]*Cell, 0, 5)
	for _, c := range cells {
		if c.Equals(*val) {
			continue
		}
		newCells = append(newCells, c)
	}
	return newCells
}

func (c *Cell) Link(other *Cell, bidi bool) {
	if other == nil {
		return
	}
	c.Links = append(c.Links, other)
	if bidi == true {
		other.Link(c, false)
	}
}

func (c *Cell) Unlink(other *Cell) {
	if other == nil {
		return
	}
	c.Links = remove(c.Links, other)
}

func (c *Cell) IsLinked(other *Cell) bool {
	if other == nil {
		return false
	}
	return find(c.Links, other)
}

func (c *Cell) ToString() string {
	s := fmt.Sprintf("x=%d,y=%d,n=%p,s=%p,e=%p,w=%p", c.x, c.y, c.North, c.South, c.East, c.West)
	return s
}

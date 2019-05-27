package basics

import (
	"bytes"
	"fmt"
	"strings"
)

const zero = int8(0)
const one = int8(1)

type Grid struct {
	rows  int8
	cols  int8
	cells [][]*Cell
}

func BuildGrid(rows int8, cols int8) *Grid {
	if rows <= zero {
		return nil
	}
	if cols <= zero {
		return nil
	}
	var cells = make([][]*Cell, rows)
	for i := zero; i < rows; i++ {
		cells[i] = make([]*Cell, cols)
		for j := zero; j < cols; j++ {
			cells[i][j] = BuildCell(i, j)
		}
	}
	g := Grid{rows, cols, cells}
	g.config()
	return &g
}

func (g *Grid) get(x int8, y int8) *Cell {
	if x < 0 {
		return nil
	}
	if y < 0 {
		return nil
	}
	if x > (g.rows - one) {
		return nil
	}
	if y > (g.cols - one) {
		return nil
	}
	return g.cells[x][y]
}

func (g *Grid) Size() int16 {
	return int16(g.rows) * int16(g.cols)
}

func (g *Grid) EachCell() chan *Cell {
	outputChannel := make(chan *Cell)
	go func() {
		for i := zero; i < g.rows; i++ {
			for j := zero; j < g.cols; j++ {
				outputChannel <- g.cells[i][j]
			}
		}
		close(outputChannel)
	}()
	return outputChannel
}

func (g *Grid) config() {
	for i := zero; i < g.rows; i++ {
		for j := zero; j < g.cols; j++ {
			g.cells[i][j].East = g.get(i, j+one)
			g.cells[i][j].West = g.get(i, j-one)
			g.cells[i][j].North = g.get(i-one, j)
			g.cells[i][j].South = g.get(i+one, j)
		}
	}
}

func (g *Grid) ToString() string {
	var output bytes.Buffer
	for i, r := range g.cells {
		output.WriteString(fmt.Sprintf("\n%d:\t", i))
		for _, c := range r {
			output.WriteString(fmt.Sprintf("%s\n", c.ToString()))
		}
	}
	return output.String()
}

func (g *Grid) PrintGrid() string {
	var output bytes.Buffer
	output.WriteString("+")
	output.WriteString(strings.Repeat("---+", int(g.cols)))
	output.WriteString("\n")
	for i := zero; i < g.rows; i++ {
		var top bytes.Buffer
		var bottom bytes.Buffer
		top.WriteString("|")
		bottom.WriteString("+")
		for j := zero; j < g.cols; j++ {
			top.WriteString(g.cells[i][j].Value)
			if g.cells[i][j].isLinked(g.cells[i][j].East) {
				top.WriteString(" ")
			} else {
				top.WriteString("|")
			}
			if g.cells[i][j].isLinked(g.cells[i][j].South) {
				bottom.WriteString("   ")
			} else {
				bottom.WriteString("---")
			}
			bottom.WriteString("+")
		}
		output.WriteString(top.String())
		output.WriteString("\n")
		output.WriteString(bottom.String())
		output.WriteString("\n")
	}
	return output.String()
}

func (g *Grid) GetCell(r int8, c int8) *Cell {
	return g.cells[r][c]
}

func (g *Grid) ClearValues() {
	for c := range g.EachCell() {
		c.Value = "   "
	}
}

func (g *Grid) GetRandomCell() *Cell {
	row := GetRand8(g.rows)
	col := GetRand8(g.cols)
	return g.GetCell(row, col)
}

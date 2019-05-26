package basics

import (
	"strings"
	"testing"
)

func TestEqualsNotEquals(t *testing.T) {
	var c1 = BuildCell(1, 2)
	var c2 = BuildCell(2, 2)
	if c1.equals(*c2) == true {
		t.Error("Error: should not be equal...", c1, c2)
	}
}

func TestEqualsEquals(t *testing.T) {
	var c1 = BuildCell(1, 2)
	var c2 = BuildCell(1, 2)
	if c1.equals(*c2) == false {
		t.Error("Error: should be equal...", c1, c2)
	}
}

func TestFindCellThatIsThere(t *testing.T) {
	var cells = createCellSlice()
	if find(cells, BuildCell(1, 2)) == false {
		t.Error("Error: should have found cell")
	}
}

func TestFindCellThatIsNotThere(t *testing.T) {
	var cells = createCellSlice()
	if find(cells, BuildCell(1, 1)) == true {
		t.Error("Error: should not have found cell")
	}
}

func TestFindCellWithNilList(t *testing.T) {
	if find(nil, BuildCell(1, 1)) == true {
		t.Error("Error: should not have found cell")
	}
}

func TestFindCellThatIsNil(t *testing.T) {
	var cells = createCellSlice()
	if find(cells, nil) == true {
		t.Error("Error: should not have found cell")
	}
}

func TestRemoveCellValid(t *testing.T) {
	cells := createCellSlice()
	if len(cells) != 4 {
		t.Error("Should start with four cells")
	}
	var newCells = remove(cells, BuildCell(1, 2))
	if find(newCells, BuildCell(1, 2)) == true {
		t.Error("Should have removed cell")
	}
	if len(newCells) != 3 {
		t.Error("Should only have three cells")
	}
}
func TestRemoveCellInvalid(t *testing.T) {
	cells := createCellSlice()
	if len(cells) != 4 {
		t.Error("Should start with four cells")
	}
	var newCells = remove(cells, BuildCell(1, 1))
	if len(newCells) != 4 {
		t.Error("Should end with four cells")
	}
}

func TestLinkCells(t *testing.T) {
	cells := createCellSlice()
	var c1 = cells[1]
	var c2 = cells[2]
	c1.Link(c2, false)
	if find(c1.Links, c2) == false {
		t.Error("c1 should be linked to c2")
	}
	if find(c2.Links, c1) == true {
		t.Error("c2 should not be linked to c1")
	}
	c1.Unlink(c2)
	if find(c1.Links, c2) == true {
		t.Error("c1 should not be linked to c2")
	}
}

func TestLinkCellsBi(t *testing.T) {
	cells := createCellSlice()
	var c1 = cells[1]
	var c2 = cells[2]
	c1.Link(c2, true)
	if find(c1.Links, c2) == false {
		t.Error("c1 should be linked to c2")
	}
	if find(c2.Links, c1) == false {
		t.Error("c2 should be linked to c1")
	}
	c1.Unlink(c2)
	if find(c1.Links, c2) == true {
		t.Error("c1 should not be linked to c2")
	}
	if c1.isLinked(c2) == true {
		t.Error("c1 should not be linked to c2")
	}
	if c1.isLinked(nil) == true {
		t.Error("c1 cannot be linked to nil")
	}
}

func TestLinkCellsNil(t *testing.T) {
	cells := createCellSlice()
	var c1 = cells[1]
	c1.Link(nil, true)
	if len(c1.Links) != 0 {
		t.Error("c1 should be have no links")
	}
	c1.Unlink(nil)
}

func TestToString(t *testing.T) {
	c1 := BuildCell(0, 0)
	s := c1.ToString()
	if strings.Compare(s, "x=0,y=0,n=0x0,s=0x0,e=0x0,w=0x0") != 0 {
		t.Errorf("bad string %s", s)
	}
}

func createCellSlice() []*Cell {
	var cells = make([]*Cell, 0, 5)

	cells = append(cells, BuildCell(1, 2))
	cells = append(cells, BuildCell(3, 2))
	cells = append(cells, BuildCell(2, 2))
	cells = append(cells, BuildCell(1, 3))
	return cells
}

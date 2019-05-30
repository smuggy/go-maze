package builders

import (
	"github.com/smuggy/go-maze/internal/basics"
	"testing"
)

func TestNilCreation(t *testing.T) {
	var g *basics.Grid = nil
	sw := Sidewinder{}
	sw.Build(g)
}

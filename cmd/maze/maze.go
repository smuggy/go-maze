package main

import (
	"fmt"
	"github.com/smuggy/go-maze/internal/basics"
	"github.com/smuggy/go-maze/internal/builders"
	"github.com/smuggy/go-maze/internal/solvers"
	"os"
	"strconv"
)

func sum(x int, y int) int {
	return x + y
}

func usage(programName string) {
	fmt.Printf("Usage: %s rows cols\n\n", programName)
	os.Exit(1)
}

func main() {
	fmt.Print("Starting application...\n\n")
	var rows int8 = 5
	var cols int8 = 5
	if len(os.Args) == 3 {
		r, e := strconv.ParseInt(os.Args[1], 10, 8)
		if e == nil {
			rows = int8(r)
		} else {
			usage(os.Args[0])
		}
		c, e := strconv.ParseInt(os.Args[2], 10, 8)
		if e == nil {
			cols = int8(c)
		} else {
			usage(os.Args[0])
		}
	} else if len(os.Args) != 1 {
		usage(os.Args[0])
	}
	//	var d = basics.BuildCell(2, 1)
	//	fmt.Println(d)
	var g = basics.BuildGrid(rows, cols)
	var b = createBuilder(builders.AldousAlgo)
	b.Build(g)
	fmt.Print(g.PrintGrid())
	//	fmt.Println("\nNext builder...")
	//	var g2 = basics.BuildGrid(4, 5)
	//	builders.Sidewinder(g2)
	//	fmt.Println(g2.PrintGrid())
	fmt.Println("\nApply solver to grid...")

	var solution solvers.Solution = solvers.DijkstraSolver(g.GetCell(0, 0))
	solution.ApplyDistances()
	fmt.Println(g.PrintGrid())

	g.ClearValues()
	fmt.Println("Find maximum path...")
	ms := solution.MaxPath()
	ms.ApplyDistances()
	fmt.Println(g.PrintGrid())

	fmt.Println("\nLeaving application...")
}

func createBuilder(builder builders.Builder) builders.MazeBuilder {
	switch builder {
	case builders.AldousAlgo:
		return builders.Aldous{}
	case builders.BinaryTreeAlgo:
		return builders.BinaryTree{}
	case builders.SidewinderAlgo:
		return builders.Sidewinder{}
	}
	return builders.BinaryTree{}
}

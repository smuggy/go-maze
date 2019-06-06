package main

import (
	"flag"
	"fmt"
	"github.com/smuggy/go-maze/internal/basics"
	"github.com/smuggy/go-maze/internal/builders"
	"github.com/smuggy/go-maze/internal/solvers"
	"os"
	"strconv"
)

var builderAlgorithm = builders.AldousAlgo
var printFullMaze = false
var printEmptyMaze = false

func sum(x int, y int) int {
	return x + y
}

func usage(programName string) {
	fmt.Printf("Usage: %s [--algo=algorithm] [--empty] [--full] rows cols\n\n", programName)
	os.Exit(1)
}

func processArgs() (int8, int8) {
	var rows int8 = 5
	var cols int8 = 5
	var progName = os.Args[0]

	algo := flag.String("algo", "aldous", "Maze building algorithm")
	flag.BoolVar(&printEmptyMaze, "empty", false, "Print empty maze")
	flag.BoolVar(&printFullMaze, "full", false, "Print full maze")
	flag.Parse()

	switch *algo {
	case "aldous":
		builderAlgorithm = builders.AldousAlgo
	case "binary":
		builderAlgorithm = builders.BinaryTreeAlgo
	case "sidewinder":
		builderAlgorithm = builders.SidewinderAlgo
	default:
		builderAlgorithm = builders.SidewinderAlgo
	}

	args := flag.Args()
	if len(args) == 2 {
		r, e := strconv.ParseInt(args[0], 10, 8)
		if e == nil {
			rows = int8(r)
		} else {
			usage(progName)
		}
		c, e := strconv.ParseInt(args[1], 10, 8)
		if e == nil {
			cols = int8(c)
		} else {
			usage(progName)
		}
	} else if len(args) != 0 {
		usage(progName)
	}
	return rows, cols
}

func main() {
	fmt.Print("Starting application...\n\n")
	rows, cols := processArgs()
	var g = basics.BuildGrid(rows, cols)
	var b = createBuilder(builderAlgorithm)
	b.Build(g)
	if printEmptyMaze {
		fmt.Print(g.PrintGrid())
	}
	fmt.Println("\nApply solver to grid...")

	var solution solvers.Solution = solvers.DijkstraSolver(g.GetCell(0, 0))
	if printFullMaze {
		solution.ApplyDistances()
		fmt.Println(g.PrintGrid())
	}

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

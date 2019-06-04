package main

import (
	"fmt"
	"github.com/smuggy/go-maze/internal/basics"
	"github.com/smuggy/go-maze/internal/builders"
	"github.com/smuggy/go-maze/internal/solvers"
)

func sum(x int, y int) int {
	return x + y
}
func main() {
	fmt.Print("Starting application...\n\n")
	//	var d = basics.BuildCell(2, 1)
	//	fmt.Println(d)
	var g = basics.BuildGrid(10, 12)
	var b = createBuilder()
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

	fmt.Println("\nLeaving application...")
}
func createBuilder() builders.MazeBuilder {
	//	return builders.BinaryTree{}
	//	return builders.Sidewinder{}
	return builders.Aldous{}
}

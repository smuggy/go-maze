package main

import (
	"basics"
	"builders"
	"fmt"
)

func sum(x int, y int) int {
	return x + y
}
func main() {
	fmt.Print("Starting application...\n\n")
	//	var d = basics.BuildCell(2, 1)
	//	fmt.Println(d)
	var g = basics.BuildGrid(4, 4)
	builders.BuildMaze(g)
	fmt.Print(g.PrintGrid())
	fmt.Println("\nLeaving application...")
}

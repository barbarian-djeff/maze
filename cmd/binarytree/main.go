package main

import (
	"log"

	"github.com/barbarian-djeff/maze/pkg/grid"
)

func main() {
	g := grid.NewGrid(10, 10)
	log.Println("start binary tree generation with size", g.Size())
	grid.Display(g)

	// apply algo
}

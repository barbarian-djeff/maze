package main

import (
	"log"
	"github.com/barbarian-djeff/maze/pkg/grid"
)

func main() {
	grid := grid.NewGrid(10, 10)
	log.Println("start binary tree generation with size", grid.Size())
}

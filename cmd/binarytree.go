package main

import (
	"log"
	"github.com/barbarian-djeff/maze/pkg/maze"
)

func main() {
	grid := maze.NewGrid(10, 10)
	log.Println("start binary tree generation with size", grid.Size())
}

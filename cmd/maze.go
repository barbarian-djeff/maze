package cmd

import "log"

type Grid struct {
	Row, Column int32
}

type Position struct {
	Row, Column int32
}

type Cell struct {
	Pos                      Position
	North, South, East, West Position
}

func (grid *Grid) Size() int32 {
	return grid.Column * grid.Row
}

func newGrid(row int32, column int32) *Grid {
	return &Grid{
		Row:    row,
		Column: column,
	}
}

func main() {
	grid := newGrid(10, 10)
	log.Printf("start maze generation with size", grid.Size())
}

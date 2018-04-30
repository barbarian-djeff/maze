package grid

type Grid struct {
	Row, Column int
	Cells       [][]Cell
}

type Cell struct {
	Up, Left, Right, Down *Cell
}

func (grid *Grid) Size() int {
	return grid.Column * grid.Row
}

// create cells without neighbours
func NewGrid(nrow, ncol int) Grid {
	cells := make([][]Cell, nrow)
	for r := 0; r < nrow; r++ {
		cells[r] = make([]Cell, ncol)
		for c := 0; c < ncol; c++ {
			cells[r][c] = Cell{}
		}
	}
	return Grid{
		Row:    nrow,
		Column: ncol,
		Cells:  cells,
	}
}

package maze

type Grid struct {
	Row, Column int
	Cells       map[int]Cell
}

type Position struct {
	Row, Column int
}

type Cell struct {
	Pos                      Position
	North, South, East, West bool
}

func (grid *Grid) Size() int {
	return grid.Column * grid.Row
}

// create cells without neighbours
func NewGrid(row int, column int) Grid {
	var cells = map[int]Cell{}
	for r := 0; r < row; r++ {
		for c := 0; c < column; c++ {
			pos := Position{
				Row:    r,
				Column: c,
			}
			hash := hash(pos)
			cells[hash] = Cell{
				Pos:   pos,
				North: false,
				South: false,
				West:  false,
				East:  false,
			}
		}
	}
	return Grid{
		Row:    row,
		Column: column,
		Cells:  cells,
	}
}

func hash(pos Position) int {
	return pos.Column*37 + pos.Row
}


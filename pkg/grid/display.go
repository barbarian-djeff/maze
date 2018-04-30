package grid

import "fmt"

func Display(g Grid) {
	var lastLine *string
	for r, row := range g.Cells {
		border := "+"
		for _, cell := range row {
			if cell.Up != nil {
				border += "   +"
			} else {
				border += "---+"
			}
		}
		if r == 0 {
			lastLine = &border
		}
		fmt.Println(border)

		line := "|"
		for _, cell := range row {
			if cell.Right != nil {
				line += "   "
			} else {
				line += "   |"
			}
		}
		fmt.Println(line)
	}
	if lastLine != nil {
		println(*lastLine)
	}
}

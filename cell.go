package life

type Cell struct {
	x int
	y int
}

type Cells []Cell

func (cell Cell) Neighbors() []Cell {
	return []Cell{
		cell.translate(-1, -1),
		cell.translate(-1, 0),
		cell.translate(-1, 1),
		cell.translate(0, -1),
		cell.translate(0, 1),
		cell.translate(1, -1),
		cell.translate(1, 0),
		cell.translate(1, 1),
	}
}

func (c Cell) translate(x int, y int) Cell {
	return Cell{c.x + x, c.y + y}
}

// sort.Interface
func (cells Cells) Less(i, j int) bool {
	if cells[i].x < cells[j].x {
		return true
	}
	return (cells[i].x == cells[j].x && cells[i].y < cells[j].y)
}

func (cells Cells) Len() int {
	return len(cells)
}

func (cells Cells) Swap(i, j int) {
	cells[i], cells[j] = cells[j], cells[i]
}

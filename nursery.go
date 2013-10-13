package life

type Nursery map[Cell]NurseryCell

func (nursery Nursery) MarkCells(cells Cells) {
	for _, cell := range cells {
		nursery.markCell(cell)
	}
}

func (nursery Nursery) markCell(cell Cell) {
	for _, neighbor := range cell.Neighbors() {
		nursery[neighbor] = nursery[neighbor].AddNeighbor()
	}
	nursery[cell] = nursery[cell].MarkAlive()
}

func (nursery Nursery) NextGeneration() Cells {
	var next = Cells{}
	for cell, spawn := range nursery {
		if spawn.Alive() {
			next = append(next, cell)
		}
	}
  return next
}


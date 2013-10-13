package life

type NurseryCell struct {
	wasAlive  bool
	neighbors int
}

func (nc NurseryCell) AddNeighbor() NurseryCell {
  return NurseryCell{nc.wasAlive, nc.neighbors + 1}
}

func (nc NurseryCell) MarkAlive() NurseryCell {
  return NurseryCell{true, nc.neighbors}
}

func (nc NurseryCell) Alive() bool {
	if nc.neighbors == 2 {
		return nc.wasAlive
	}

	return nc.neighbors == 3
}

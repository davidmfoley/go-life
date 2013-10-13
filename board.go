package life

type Board struct {
  cells Cells
}

func (board Board) Evolve() Board {
  var nursery = board.buildNursery()
  return Board{nursery.NextGeneration()}
}

func (board Board) buildNursery() Nursery {
  var nursery = Nursery{}
  nursery.MarkCells(board.Cells())
  return nursery
}

func (board Board) Cells() Cells {
  return board.cells
}

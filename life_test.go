package life

import "testing"
import "fmt"
import "sort"

func TestNeighbors(t *testing.T) {
  var neighbors = Cell{0, 0}.Neighbors()
  var expected = []Cell{
    Cell{-1, -1},
    Cell{-1, 0},
    Cell{-1, 1},
    Cell{0, -1},
    Cell{0, 1},
    Cell{1, -1},
    Cell{1, 0},
    Cell{1, 1},
  }
  AssertCellsEqual(expected, neighbors, t)
}

func TestOscillator(t *testing.T) {
  var starting = Cells{
    Cell{0, 0},
    Cell{0, 1},
    Cell{0, 2},
  }
  var expected = Cells{
    Cell{1, 1},
    Cell{-1, 1},
    Cell{0, 1},
  }

  CheckEvolution(t, starting, expected)
  CheckEvolution(t, expected, starting)
}

func TestBox(t *testing.T) {
  var box = Cells{
    Cell{23, 26},
    Cell{24, 26},
    Cell{23, 27},
    Cell{24, 27},
  }

  CheckEvolution(t, box, box)
}

func TestDiamond(t *testing.T) {
  var box = Cells{
    Cell{23, 26},
    Cell{24, 25},
    Cell{25, 26},
    Cell{24, 27},
  }

  CheckEvolution(t, box, box)
}

func CheckEvolution(t *testing.T, starting Cells, expected Cells) {
  var evolved = Board{starting}.Evolve()
  AssertCellsEqual(expected, evolved.Cells(), t)
}

func AssertCellsEqual(expected Cells, actual Cells, t *testing.T) {
  sort.Sort(expected)
  sort.Sort(actual)
  AssertEqual(expected, actual, t)
}

func AssertEqual(expected interface{}, actual interface{}, t *testing.T) {
  var expectedFormat = fmt.Sprintf("%v", expected)
  var actualFormat = fmt.Sprintf("%v", actual)

  if expectedFormat != actualFormat {
    t.Errorf("expected %v but got %v", expectedFormat, actualFormat)
  }
}


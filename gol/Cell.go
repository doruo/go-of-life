package gol

type Cell struct {
	I, J      int
	isAlive   bool
	Adjacents []Cell
}

// Creates a new cell with alive adjacents in memory.
func NewCell(i, j int) *Cell {
	return &Cell{
		I:         i,
		J:         j,
		isAlive:   false,
		Adjacents: make([]Cell, 8), // All cell adjacent to this
	}
}

// --------------------------------------------

// Update cell state based on Conway's rules
func (c *Cell) UpdateState() {
	c.isAlive = c.GetUpdatedState()
}

// --------------------------------------------

// Returns updated cell state
// based on Conway's rules.
func (c *Cell) GetUpdatedState() bool {
	// 1 - Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	if c.isAlive && len(c.Adjacents) < 2 {
		return false
	}
	// 2 - Any live cell with two or three live neighbours lives on to the next generation.
	if c.isAlive && (len(c.Adjacents) == 2 || len(c.Adjacents) == 3) {
		return true
	}
	// 3 - Any live cell with more than three live neighbours dies, as if by overpopulation.
	if c.isAlive && len(c.Adjacents) > 3 {
		return false
	}
	// 4 - Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if !c.isAlive && len(c.Adjacents) == 3 {
		return true
	}
	return false
}

func (c *Cell) IsAdjacent(c2 Cell) bool {
	for _, cell := range c.Adjacents {
		if &cell == &c2 {
			return true
		}
	}
	return false
}

func (c *Cell) IsAlive() bool {
	return c.isAlive
}

func (c *Cell) SetAlive(isAlive bool) {
	c.isAlive = isAlive
}

func (c *Cell) SetAdjacents(adjs []Cell) {
	c.Adjacents = adjs
}

func (c *Cell) SetCoordinates(i, j int) {
	c.I, c.J = i, j
}

// --------------------------------------------

func (c *Cell) ToString() string {
	if c.isAlive {
		return (Green() + "O" + ColorReset())
	}
	return (Red() + "~" + ColorReset())
}

package minesweeper

type cell struct {
	y, x      int
	IsMine    bool
	IsFlag    bool
	Neighbors int
	IsActive  bool
}

func createCell(y, x int) *cell {
	return &cell{y, x, false, false, 0, false}
}

func (c *cell) toggleFlag() {
	c.IsFlag = !c.IsFlag
}

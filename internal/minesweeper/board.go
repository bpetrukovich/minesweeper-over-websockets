package minesweeper

import (
	"math/rand"
)

type board struct {
	Height, Width int
	Grid          [][]cell
	Mines         int
	actives       int
}

func makeBoard(height, width int, mines int) *board {
	grid := make([][]cell, height)
	for row := range grid {
		grid[row] = make([]cell, width)
		for col := range grid[row] {
			grid[row][col] = *createCell(row, col)
		}
	}

	b := &board{height, width, grid, mines, 0}
	b.generateMines()
	return b
}

func (b *board) checkX(x int) bool {
	return x >= 0 && x < b.Width
}

func (b *board) checkY(y int) bool {
	return y >= 0 && y < b.Height
}

func (b *board) checkCoords(y, x int) bool {
	return b.checkX(x) && b.checkY(y)
}

func (b *board) generateMines() {
	minesGenerated := 0
	for {
		y := rand.Intn(b.Height)
		x := rand.Intn(b.Width)
		if !b.Grid[y][x].IsMine {
			b.Grid[y][x].IsMine = true
			minesGenerated++

			nbs := b.getNeighbors(&b.Grid[y][x])
			for _, nb := range nbs {
				nb.Neighbors++
			}
		}
		if minesGenerated == b.Mines {
			break
		}
	}
}

func (b *board) getNeighbors(c *cell) []*cell {
	cells := []*cell{}
	if b.checkCoords(c.y-1, c.x-1) {
		cells = append(cells, &b.Grid[c.y-1][c.x-1])
	}
	if b.checkCoords(c.y-1, c.x+1) {
		cells = append(cells, &b.Grid[c.y-1][c.x+1])
	}
	if b.checkCoords(c.y+1, c.x-1) {
		cells = append(cells, &b.Grid[c.y+1][c.x-1])
	}
	if b.checkCoords(c.y+1, c.x+1) {
		cells = append(cells, &b.Grid[c.y+1][c.x+1])
	}
	if b.checkCoords(c.y-1, c.x) {
		cells = append(cells, &b.Grid[c.y-1][c.x])
	}
	if b.checkCoords(c.y, c.x+1) {
		cells = append(cells, &b.Grid[c.y][c.x+1])
	}
	if b.checkCoords(c.y+1, c.x) {
		cells = append(cells, &b.Grid[c.y+1][c.x])
	}
	if b.checkCoords(c.y, c.x-1) {
		cells = append(cells, &b.Grid[c.y][c.x-1])
	}

	return cells
}

func (b *board) setActive(c *cell) {
	if !c.IsActive && !c.IsMine && !c.IsFlag {
		c.IsActive = true
		b.actives++
	}
}

func (b *board) open(c *cell) {
	b.setActive(c)
	if c.Neighbors == 0 {
		queue := []*cell{c}
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			b.setActive(cur)

			nbs := b.getNeighbors(cur)
			if cur.Neighbors == 0 {
				for _, nb := range nbs {
					if !nb.IsActive {
						queue = append(queue, nb)
					}
				}
			}
		}
	}
}

func (b *board) toggleFlag(y, x int) {
	b.Grid[y][x].toggleFlag()
}

// func (b *board) tick() {
// 	fmt.Print("\033[H\033[2J")
// 	fmt.Println(b.actives, b.width*b.height-b.mines)
// 	b.Print()
// 	if b.gameOver {
// 		fmt.Println("Game Over!!!")
// 		return
// 	}
// 	if b.actives == b.width*b.height-b.mines {
// 		fmt.Println("VIN!!!")
// 		return
// 	}
// }

// func (b *board) Interact(key rune) bool {
// 	defer b.tick()
// 	switch key {
// 	case '\x00':
// 		b.press(b.y, b.x)
// 		if b.gameOver || (b.actives == b.width*b.height-b.mines) {
// 			return false
// 		}
// 	case 'w':
// 		b.selectCell(b.y-1, b.x)
// 	case 'a':
// 		b.selectCell(b.y, b.x-1)
// 	case 's':
// 		b.selectCell(b.y+1, b.x)
// 	case 'd':
// 		b.selectCell(b.y, b.x+1)
// 	case 'p':
// 		b.grid[b.y][b.x].toggleFlag()
// 	case 'q':
// 		b.gameOver = true
// 		return false
// 	}
// 	return true
// }

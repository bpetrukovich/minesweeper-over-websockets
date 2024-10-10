package minesweeper

type game struct {
	Board    *board
	selected map[*Player]*cell
	GameOver bool
	Win      bool
}

func makeGame(height, width, mines int) *game {
	return &game{
		Board:    makeBoard(height, width, mines),
		selected: make(map[*Player]*cell),
		GameOver: false,
		Win:      false,
	}
}

func (g *game) outData() *msgOut {
	action := "play"
	if g.GameOver {
		action = "gameOver"
	} else if g.Win {
		action = "win"
	}
	return &msgOut{
		Action: action,
		Board:  g.Board,
	}
}

func (g *game) press(y, x int) {
	b := g.Board
	c := &b.Grid[y][x]
	if c.IsFlag {
		return
	}
	if c.IsMine {
		g.GameOver = true
		return
	}
	if c.IsActive {
		nbs := b.getNeighbors(c)
		openNbs := 0
		for _, nb := range nbs {
			if nb.IsFlag && nb.IsMine {
				openNbs++
			}
		}
		if openNbs == c.Neighbors {
			for _, nb := range nbs {
				b.open(nb)
			}
		}
	}

	b.open(c)

	if b.actives == b.Width*b.Height-b.Mines {
		g.Win = true
	}
}

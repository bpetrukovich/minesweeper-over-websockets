package minesweeper

import "mine/internal/ws"

type Room struct {
	hub     *ws.Hub
	players []*Player
	G       *game
}

func MakeRoom(h *ws.Hub) *Room {
	return &Room{
		players: make([]*Player, 0),
		hub:     h,
	}
}

func (r *Room) AddPlayer(p *Player) {
	r.players = append(r.players, p)
}

func (r *Room) Join(client *ws.Client) {
	p := &Player{client: client}
	r.players = append(r.players, p)
}

func (r *Room) newGame(height, width, mines int) {
	r.G = makeGame(height, width, mines)
}

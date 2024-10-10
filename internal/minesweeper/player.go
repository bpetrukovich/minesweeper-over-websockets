package minesweeper

import (
	"encoding/json"
	"log"
	"mine/internal/ws"
)

type Player struct {
	room   *Room
	client *ws.Client
}

func MakePlayer(r *Room, c *ws.Client) *Player {
	p := &Player{
		room:   r,
		client: c,
	}
	r.AddPlayer(p)
	return p
}

type msgIn struct {
	Action string `json:"action"`
	C      coords `json:"coords"`
}

type msgOut struct {
	Action string `json:"action"`
	Board  *board `json:"board"`
}

type coords struct {
	Y int `json:"y"`
	X int `json:"x"`
}

func (p *Player) Play() {
	in := p.client.Get
	// out := p.client.Send
	for {
		message := <-in
		log.Printf("string(message): %v\n", string(message))
		m := msgIn{}
		err := json.Unmarshal(message, &m)
		if err != nil {
			log.Printf("cant unmarshal message: %v", err)
			continue
		}
		g := p.room.G
		switch m.Action {
		case "click":
			g.press(m.C.Y, m.C.X)
		case "toggleFlag":
			g.Board.toggleFlag(m.C.Y, m.C.X)
		case "newGame":
			p.room.newGame(16, 16, 40)
			g = p.room.G
		}

		j, err := json.Marshal(g.outData())
		if err != nil {
			log.Printf("cant serialize to json")
			return
		}
		p.room.hub.Broadcast <- j
	}

}

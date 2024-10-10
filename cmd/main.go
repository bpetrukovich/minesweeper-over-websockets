package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mine/internal/minesweeper"
	"mine/internal/ws"
	"net/http"

	"github.com/gorilla/websocket"
)

var publicHandler http.Handler = http.FileServer(http.Dir("../public"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "public/main.html", http.StatusFound)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request, h *ws.Hub, room *minesweeper.Room) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	c := ws.MakeClient(h, conn)
	h.AddClient(c)
	room.Join(c)
	p := minesweeper.MakePlayer(room, c)
	_ = p

	j, err := json.Marshal(room.G)
	if err != nil {
		fmt.Printf("cant serialize to json")
		return
	}
	h.Broadcast <- j

	go p.Play()
	go c.Read()
	go c.Write()
}

func main() {
	hub := ws.MakeHub()
	go hub.Run()

	globalRoom := minesweeper.MakeRoom(hub)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(w, r, hub, globalRoom)
	})
	http.Handle("/public/", http.StripPrefix("/public/", publicHandler))
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}

package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	h *Hub

	conn *websocket.Conn

	Send chan []byte
	Get  chan []byte
}

func MakeClient(h *Hub, conn *websocket.Conn) *Client {
	return &Client{
		h:    h,
		conn: conn,
		Send: make(chan []byte, 256),
		Get:  make(chan []byte, 256),
	}
}

func (c *Client) Read() {
	defer func() {
		c.h.removeClient(c)
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Println("read:", string(message))
		c.Get <- message
	}
}

func (c *Client) Write() {
	defer func() {
		c.h.removeClient(c)
		c.conn.Close()
	}()
	for {
		message := <-c.Send
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write:", err)
		}
	}
}

func (c *Client) SendMessage(message []byte) {
	c.Send <- message
}

func (c *Client) getMessage() []byte {
	message := <-c.Get
	return message
}

package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type WsContext struct {
	conn    *websocket.Conn
	reader  chan *[]byte
	spoiler chan error
}

func (c *WsContext) start() {
	c.reader = make(chan *[]byte)
	c.spoiler = make(chan error)
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	go c.read()
	go c.startPingPing()
}

func (c *WsContext) startPingPing() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *WsContext) WriteMessage(message string) {
	log.Println("Writing =>", message)
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println(err.Error())
		c.spoiler <- err
	}
}

func (c *WsContext) WriteBinrayData(data *[]byte) {

}

func (c *WsContext) read() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			c.spoiler <- err
			return
		}
		c.reader <- &message
	}
}

func (c *WsContext) Close() {
	log.Println("Closing connection")
	c.conn.WriteMessage(websocket.CloseMessage, []byte{})
	c.conn.Close()
	close(c.reader)
	close(c.spoiler)
}

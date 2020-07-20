package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	closed   chan bool
	conn     *websocket.Conn
	Response chan string
}

func NewClient() *Client {
	return &Client{
		closed:   make(chan bool, 1),
		Response: make(chan string),
	}
}

func (c *Client) main() {
	go func() {
		for {
			_, buffer, err := c.conn.ReadMessage()
			if err != nil {
				log.Printf("Error when reading from connection: %s", err)
			}
			c.Response <- string(buffer)
			select {
			case <-c.closed:
				return
			default:
				continue
			}
		}
	}()
}

func (c *Client) Connect(url string) error {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	c.main()

	return nil
}

func (c *Client) Subscribe(s string) error {
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(s))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() error {
	defer c.conn.Close()
	c.closed <- true
	err := c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)
	if err != nil {
		return err
	}
	return nil
}

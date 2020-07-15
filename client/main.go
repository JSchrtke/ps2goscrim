package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	eventStream := url.URL{
		Scheme:   "wss",
		Host:     "push.planetside2.com",
		Path:     "streaming",
		RawQuery: "environment=ps2&service-id=s:ps2goscrim",
	}

	log.Printf("connecting to %s", eventStream.String())
	connection, _, err := websocket.DefaultDialer.Dial(eventStream.String(), nil)
	if err != nil {
		log.Fatal("Error when dialing: ", err)
	}
	defer connection.Close()

	go func() {
		for {
			_, buffer, err := connection.ReadMessage()
			if err != nil {
				log.Println("Error when reading message: ", err)
			}
			log.Printf("Recieved message: %s", buffer)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			log.Println("closing connection")
			err := connection.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)
			if err != nil {
				log.Println("Error when closing connection: ", err)
				return
			}
			return
		}
	}
}

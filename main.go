package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	client := NewClient()
	err := client.Connect("wss://push.planetside2.com/streaming?environment=ps2&service-id=s:ps2goscrim")
	if err != nil {
		log.Printf("Error when connecting: %s", err)
	}

	id := "8267848801130180513"
	s := fmt.Sprintf(
		"{ \"service\":\"event\","+
			"\"action\":\"subscribe\","+
			"\"characters\":[\"%s\"],"+
			"\"eventNames\":[\"Death\"] }",
		id,
	)
	err = client.Subscribe(s)
	if err != nil {
		log.Printf("Error when subscribing to: %s", s)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			err = client.Close()
			if err != nil {
				panic(err)
			}
			return
		case res := <-client.Response:
			log.Printf("Got response: %s", res)
		}
	}
}

package main

import (
	"fmt"

	"github.com/r3labs/sse/v2"
)

func main() {
	// This one does not work
	path := "ENTER PATH HERE"
	client := sse.NewClient(path, sse.ClientMaxBufferSize(1<<19))

	client.OnDisconnect(func(c *sse.Client) {
		fmt.Println("disconnect!")
	})

	client.Subscribe("messages", func(msg *sse.Event) {
		fmt.Println(string(msg.Data))
		fmt.Println(string(msg.Event))
	})
}

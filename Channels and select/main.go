package main

import (
	"fmt"
	"time"
)

func main() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "Channel 1. 200ms"
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for {
			message2 <- "Channel 2. 1s"
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		default:
		}
	}
}

func example1() {
	msg := make(chan string, 3)

	msg <- "Ninja channel"
	msg <- "Ninja channel 2"
	msg <- "Ninja channel 3"

	close(msg)

	for m := range msg {
		fmt.Println(m)
	}
}

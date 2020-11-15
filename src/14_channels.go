package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		messages <- "ping"
	}()

	fmt.Println("Do some stuff here")

	fmt.Println("Now waiting")
	msg := <-messages

	fmt.Println("Now free")
	fmt.Println(msg)
}

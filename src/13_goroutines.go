package main

import (
	"fmt"
	"time"
)

var running = 0

func doThing(label string, n int, ms int) {
	running++
	for i := 0; i < n; i++ {
		fmt.Printf("[%s] #%d - Doing thing\n", label, i)
		time.Sleep(time.Millisecond * time.Duration(ms))
	}

	running--
}

func main() {
	fmt.Println("Hello world")

	doThing("serial", 5, 200)

	go doThing("parallel-a", 5, 200)
	go doThing("parallel-b", 50, 50)

	time.Sleep(time.Millisecond)
	for running > 0 {
		time.Sleep(time.Millisecond * 10)
	}
}

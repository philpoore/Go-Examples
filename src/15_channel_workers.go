package main

import (
	"fmt"
	"time"
)

func worker(i int, done chan bool) {
	fmt.Printf("[%d] working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("[%d] done\n", i)

	done <- true
}

func main() {
	done := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		go worker(i, done)
	}

	for i := 0; i < 5; i++ {
		<-done
	}
}

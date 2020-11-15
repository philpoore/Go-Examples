package main

import (
	"fmt"
	"time"
)

/*
	Simple Fib function
*/
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

var fibCache = make([]int, 50)

func fibDp(n int) int {
	if fibCache[n] > 0 {
		return fibCache[n]
	}
	if n <= 2 {
		return 1
	}
	res := fibDp(n-2) + fibDp(n-1)

	fibCache[n] = res
	return res
}

var profiles = make(map[string]time.Time)

func profileStart(key string) {
	profiles[key] = time.Now()
}

func profileEnd(key string) {
	now := time.Now()
	then := profiles[key]
	diff := (now.UnixNano() - then.UnixNano()) / 1000000
	fmt.Printf("[profile:%s] %dms\n", key, diff)
}

func main() {
	profileStart("fib")
	fmt.Println("Fib(40)", fib(42))
	profileEnd("fib")

	profileStart("fib-dp")
	fmt.Println("Fib(40) dp", fibDp(42))
	profileEnd("fib-dp")
}

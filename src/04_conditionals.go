package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("a=%d and is ", a)

	if a%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// Simple bools
	b := true
	if b {
		fmt.Println("I knew this would happen!")
	}
}

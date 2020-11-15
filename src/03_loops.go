package main

import "fmt"

func main() {
	// Simple for style loop
	for i := 0; i < 5; i++ {
		fmt.Printf("simple loop i=%d\n", i)
	}

	// While loop with test
	i := 0
	for i < 5 {
		fmt.Printf("while loop with test i=%d\n", i)
		i++
	}

	// While loop no test
	i = 0
	for {
		fmt.Printf("while loop no test i=%d\n", i)
		i++
		if i == 5 {
			break
		}
	}

}

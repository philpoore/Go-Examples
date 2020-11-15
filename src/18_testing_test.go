package main

import (
	"fmt"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("add = ", add(4, 5))
}

func TestAdd(t *testing.T) {
	var a = add(5, 10)
	if a != 15 {
		t.Errorf("TestAdd(5, 10) = %d; want 15", a)
	}
}

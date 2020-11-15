package main

import "fmt"

func main() {
	var a [10]int

	fmt.Println("Before: ", a)

	for k := range a {
		a[k] = k
	}
	for k, v := range a {
		a[k] = v * v
	}
	fmt.Println("After: ", a)

	// Sum

	items := []int{10, 67, 59, 123}
	sum := 0
	for _, v := range items {
		sum += v
	}
	fmt.Printf("sum=%d\n", sum)
}

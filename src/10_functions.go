package main

import (
	"fmt"
	"math"
)

func foo() (int, int) {
	return 10, 45
}

func sum(nums ...int) {
	total := 0
	for _, a := range nums {
		total += a
	}
	fmt.Println("sum =", total)
}

func minMax(nums ...int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	var a, b = foo()
	fmt.Printf("a=%d b=%d\n", a, b)

	sum(1, 2, 3, 4)

	c := []int{5, 6, 7, 8, 5, 4, 5, 3, 5}
	sum(c...)

	var min, max = minMax(45, 734, 64556, 34, 23, 456, 3, 234, 5, 76)
	fmt.Println("minMax", min, max)
}

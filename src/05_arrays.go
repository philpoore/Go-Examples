package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a array:", a)

	for i := 0; i < 5; i++ {
		a[i] = i
	}
	fmt.Println("a array:", a)

	fmt.Println("length of a:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// 2D array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

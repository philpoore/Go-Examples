package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	fmt.Printf("a is %d\n", a)
	fmt.Printf("b is %d\n", b)

	var c = a + b
	fmt.Printf("c is %d\n", c)

	var d = "Hi there!"

	fmt.Println(d)

	var s1 = "Hello "
	var s2 = "world"

	fmt.Println(s1 + s2)

	const foo = "Hi"
	const other = "mom!"
	fmt.Println(foo, other)
}

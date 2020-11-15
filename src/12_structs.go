package main

import "fmt"

type person struct {
	name string
	age  int
}

func reverse(input string) string {
	var output string
	for _, rune := range input {
		output = string(rune) + output
	}
	return output
}

func reverseName(p *person) {
	p.name = reverse(p.name)
}

func (p *person) display() {
	fmt.Println("Displaying", p)
}

func main() {
	phil := person{name: "phil", age: 33}
	fmt.Println(phil)

	phil.name = "Phil"
	fmt.Println(phil)

	reverseName(&phil)
	fmt.Println(phil)
	reverseName(&phil)

	phil.display()
}

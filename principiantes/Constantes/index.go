package main

import "fmt"

const (
	a = 10
	b = 20
	c = 30
	d = 40
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var fruit string = "Apple"
	fruit = "Mango"
	fmt.Println(fruit)

	// También puedes escribirlo sin especificar el tipo.
	// const manzana = "Manzana"
	// manzana = "Mango" // Error
	// fmt.Println(manzana)

	const (
		one   = "John"
		two   = "Alex"
		three = "HuXn"
	)

	// ERROR
	// one = "Alex"
	// two = "John"
	// three = "Nina"

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)

}

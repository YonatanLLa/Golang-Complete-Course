package main

import "fmt"

// var number int = 20
var number = 20

func main() {
	fmt.Println(number)

	//Operador de decoración corta
	num := 1000
	fmt.Println(num)

	// Usando nombres descriptivos
	// UPPERCASE 👎
	// lowercase 👎
	// snack_case 👎
	// CamelCase 👍
	fullName := "HuXn WebDev"
	fmt.Println(fullName)

	//Error inesperado en variables

	// 1.
	// var name string = 20 // ERROR
	// fmt.Println(name)

	// 2.
	// var var = 20
	// fmt.Println(var)

	// 3.
	// var 1student = "alex"
	// fmt.Println(1student)

	// 4.
	// var Awesome animal = "dog"
	// fmt.Println(Awesome animal)

	// 5.
	// var special = 20
	// special = "Special String"
	// fmt.Println(special)
}

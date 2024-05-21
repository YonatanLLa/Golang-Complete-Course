package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Print("Ingresa el primer numero: ")
	fmt.Scan(&number1)
	fmt.Print("Ingresa el segundo numero: ")
	fmt.Scan(&number2)

	sum := number1 + number2
	sub := number1 - number2
	mul := number1 * number2
	div := number1 / number2
	fmt.Printf("Suma: %d\nResta: %d\nMultiplicación: %d\nDivision: %d\n", sum, sub, mul, div)

}

package main

import "fmt"

func main() {
	var number int
	fmt.Print("Ingrese el numero: ")
	fmt.Scan(&number)
	count := 0
	originalNumber := number
	sum := 0
	for number != 0 {
		dig := number % 10
		fmt.Println(dig, "mod")
		sum += dig
		number /= 10
		fmt.Println(number, "div")
		count++
	}
	fmt.Printf("La suma de los dígitos del número %d es: %d\n", originalNumber, sum)
}

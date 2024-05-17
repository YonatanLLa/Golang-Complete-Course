package main

import "fmt"

// Par o Impar: Verifica si un número es par o impar.

func main() {
	num := 11
	resp := num%2 == 0

	fmt.Println("¿El número", num, "es par?", resp)
}

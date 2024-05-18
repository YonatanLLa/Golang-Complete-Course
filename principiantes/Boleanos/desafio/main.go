package main

import "fmt"

func main() {

	anio := 2100

	resp := (anio%4 == 0 && anio%100 != 0) || anio%400 == 0
	fmt.Printf("¿Es bisiesto el año %d? %t\n", anio, resp)
}

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	init := 0
	res := 0

	for init <= 10 {
		res = init * num
		fmt.Println(num, "*", init, "=", res)
		init++
	}

}

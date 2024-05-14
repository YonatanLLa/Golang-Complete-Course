# While Loop

## A diferencia de otros lenguajes de programación, Go no tiene una palabra clave dedicada para un bucle while. Sin embargo, podemos usar el bucle for para realizar la funcionalidad de un bucle while.

```go
package main
import "fmt"

func main() {
	number := 0

	for number <= 10 {
		fmt.Println(number)
		number++
	}
}
```
- En este ejemplo, el bucle for se usa para imprimir los números del 0 al 10. 
- La condición del bucle es que el valor de la variable number debe ser menor o igual a 10.
- Cada vez que se ejecuta el bucle, el valor de number se incrementa en 1.
Esto se hace mediante la instrucción number++ dentro del cuerpo del bucle.
- Cuando el valor de number es mayor que 10, el bucle se detiene y el programa termina.
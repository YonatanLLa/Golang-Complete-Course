### Ejercicios de Bucles en Go

1. **Imprimir números pares del 1 al 20**: Utiliza un bucle para imprimir todos los números pares del 1 al 20.

```go
package main
import "fmt"

func main() {
	num := 1
	for num <= 20 {
		if num%2 == 0 {
			fmt.Println(num)
		}
		num++
	}
}
```

2. **Suma de números del 1 al 100**: Utiliza un bucle para sumar todos los números del 1 al 100 e imprime el resultado.

```go
package main
import "fmt"

func main() {
	num := 1
	acc := 0
	for num <= 100 {
		acc += num
		num++
	}
	fmt.Println(acc)
}
```

3. **Tabla de multiplicar**: Escribe un programa que solicite al usuario un número y luego imprima la tabla de multiplicar de ese número del 1 al 10.

- Utiliza el siguien codigo para atrapar el numero ingresado por la terminal

```go
    var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)
```

```go
package main
import "fmt"

func main() {
	var num int
	init := 0
	res := 0

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	for init <= 10 {
		res = init * num
		fmt.Println(num, "*", init, "=", res)
		init++
	}
}
```
4. **Factorial de un número**: Escribe un programa que solicite al usuario un número y luego calcule su factorial utilizando un bucle.


5. **Generador de secuencia Fibonacci**: Escribe un programa que genere los primeros 10 números de la secuencia Fibonacci utilizando un bucle.

6. **Contador de dígitos**: Escribe un programa que cuente el número de dígitos en un número ingresado por el usuario utilizando un bucle.

7. **Suma de los dígitos de un número**: Escribe un programa que solicite al usuario un número y luego calcule la suma de sus dígitos utilizando un bucle.

8. **Búsqueda en una lista**: Escribe un programa que busque un número ingresado por el usuario en una lista predefinida utilizando un bucle.

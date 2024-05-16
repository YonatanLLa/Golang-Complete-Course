### Ejercicios de Bucles en Go

1. **Imprimir números pares del 1 al 20**: Utiliza un bucle para imprimir todos los números pares del 1 al 20.

2. **Suma de números del 1 al 100**: Utiliza un bucle para sumar todos los números del 1 al 100 e imprime el resultado.

3. **Tabla de multiplicar**: Escribe un programa que solicite al usuario un número y luego imprima la tabla de multiplicar de ese número del 1 al 10.

- Utiliza el siguien codigo para atrapar el numero ingresado por la terminal

```go
    var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)
```

4. **Factorial de un número**: Escribe un programa que solicite al usuario un número y luego calcule su factorial utilizando un bucle.

- Utiliza el siguien codigo para atrapar el numero ingresado por la terminal

```go
    var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)
```

5. **Generador de secuencia Fibonacci**: Escribe un programa que genere los primeros 10 números de la secuencia Fibonacci utilizando un bucle.

6. **Contador de dígitos**: Escribe un programa que cuente el número de dígitos en un número ingresado por el usuario utilizando un bucle.

7. **Suma de los dígitos de un número**: Escribe un programa que solicite al usuario un número y luego calcule la suma de sus dígitos utilizando un bucle.

## Solución

#### Ejercicio 1.

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

### Ejerccio 2.

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

### Ejerccio 3.

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

### Ejercicio 4.

```go
package main
import "fmt"

func main() {
	var num int
	init := 1
	acc := 1

	fmt.Print("Ingrese un numero: ")
	fmt.Scan(&num)

	for init <= num {
		acc *= init
		init++
	}
	fmt.Println("El factoril del num", num, "es: ", acc)
}
```

### Ejercicio 5.

```go
package main
import "fmt"
	// iniciando con: 0 y 1
	// 0 + 1 = 1
	// 1 + 1 = 2
	// 1 + 2 = 3
	// 2 + 3 = 5
	// 3 + 5 = 8
	// 5 + 8 = 13
	// 8 + 13 = 21
	// 13 + 21 = 34
	// Fibonacci:  0, 1, 1, 2, 3, 5, 8, 13, 21, 34

func main() {

	a := 0
	b := 1
	init := 0
	for init <= 9 {
		fmt.Println("Fibonacci", init, ": ", a)
		c := a + b
		a = b
		b = c
		init++
	}
}
```

### Ejercicio 6.

```go
package main
import "fmt"

func main() {
	var numero int
	fmt.Print("Ingresa un número: ")
	fmt.Scan(&numero)
	acc := 0
	dato := numero
	for numero != 0 {
		numero /= 10
		acc++
	}

	fmt.Printf("El número de dígitos en %d es: %d\n", dato, acc)
}
```

### Ejercicio 7.

```go
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
		sum += dig
		number /= 10
		count++
	}
	fmt.Printf("La suma de los dígitos del número %d es: %d\n", originalNumber, sum)
}
```

# Ejercicios con Operadores

### Ejercicio 1: Operadores de Incremento y Decremento

- Crear nombre de variable (lgNumber)
- Almacenar 1000 como valor
- Suma esa variable consigo misma
- Multiplica esa variable por sí misma.

### Ejercicio 2: Operadores Aritméticos

- **Descripción**: Escribe un programa en Go que tome dos números enteros como entrada y realice las siguientes operaciones: suma, resta, multiplicación, división y módulo. Imprime los resultados de cada operación.

- Con este codigo ingresas los dos numeros.

```go
var number1, number2 int
fmt.Print("Ingresa el primer numero: ")
fmt.Scan(&number1)

fmt.Print("Ingresa el segundo numero: ")
fmt.Scan(&number2)
```

### Ejercicio 3: Operadores de Comparación

- **Descripción**: Escribe un programa en Go que tome dos números enteros como entrada y compare si son iguales, diferentes, mayor, menor, mayor o igual, y menor o igual. Imprime los resultados de cada comparación.

- Con este codigo ingresas los dos numeros.

```go
var number1, number2 int
fmt.Print("Ingresa el primer numero: ")
fmt.Scan(&number1)

fmt.Print("Ingresa el segundo numero: ")
fmt.Scan(&number2)
```


### Ejercicio 4: Operadores de Asignación

- **Descripción**: Escribe un programa en Go que tome un número entero y realice operaciones de asignación con suma, resta, multiplicación y división. Imprime los resultados después de cada operación.

### Ejercicio 5: Combinación de Operadores

- **Descripción**: Escribe un programa en Go que tome dos números enteros y realice una serie de operaciones combinadas (usando aritméticos, de comparación y lógicos) para determinar si los números cumplen ciertas condiciones. Por ejemplo, verifica si la suma de los dos números es mayor que 10 y si uno de ellos es par.

### Solución

### Ejercicio 1
```go
package main
import "fmt"

func main() {
	lgNumber := 1000
	fmt.Println(lgNumber)
	fmt.Println(lgNumber + lgNumber)
	fmt.Println(lgNumber * lgNumber)
}
```

### Ejercicio 2

```go
package main
import "fmt"

func main() {
    var number1, number2 int
    fmt.Print("Ingresa el primer numero: ")
    fmt.Scan(&number1)
    fmt.Print("Ingresa el segundo numero: ")
    fmt.Scan(&number2)
    fmt.Printf("Suma de %d y %d es: %d\n", number1, number2, number1+number2)
    fmt.Printf("Multiplicacion de %d y %d es: %d\n", number1, number2, number1*number2)
}

```

### Ejercicio 3

```go
package main
import "fmt"

func main() {
	var number1, number2 int
	fmt.Print("Ingresa el primer numero: ")
	fmt.Scan(&number1)
	fmt.Print("Ingresa el segundo numero: ")
	fmt.Scan(&number2)

	fmt.Printf("Numero %d es igual a %d: %t\n", number1, number2, number1 == number2)
	fmt.Printf("Numero %d es diferente a %d: %t\n", number1, number2, number1 != number2)
	fmt.Printf("Numero %d es mayor que %d: %t\n", number1, number2, number1 > number2)
	fmt.Printf("Numero %d es menor que %d: %t\n", number1, number2, number1 < number2)
	fmt.Printf("Numero %d es mayor o igual que %d: %t\n", number1, number2, number1 >= number2)
	fmt.Printf("Numero %d es menor o igual que %d: %t\n", number1, number2, number1 <= number2)
}
```

### Ejercicio 4

```go
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

```
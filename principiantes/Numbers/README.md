# Los operadores se utilizan para realizar operaciones sobre variables y valores.

## Operadores aritméticos

### 1. Suma: El operador + suma dos operandos. Por ejemplo, x+y.
### 2. Resta: El operador - resta dos operandos. Por ejemplo, x-y.
### 3. Multiplicación: El operador * multiplica dos operandos. Por ejemplo, x*y.
### 4. División: El operador / divide el primer operando por el segundo. Por ejemplo, x/y.
### 5. Módulo: Devuelve el resto de la división.
### 6. Incremento: El ++ Incrementa el valor de una variable en 1
### 7. Decremento: Disminuye el valor de una variable en 1

```go
package main
import "fmt"

func main() {
	fmt.Println(2 + 2)  // 4
	fmt.Println(2 - 1)  // 1
	fmt.Println(2 * 2)  // 4
	fmt.Println(20 / 2) // 10
	fmt.Println(20 % 4) // 0
}
```
### Incremento
```go
package main
import "fmt"

func main() {
	number := 10
	number++ // number = number + 1
	fmt.Println(number)
}
```

### Decremento
```go
package main
import "fmt"

func main() {
    number := 10
    number-- // number = number - 1
    fmt.Println(number)
}   
```

### Operadores de Asignación

#### Los operadores de asignación se utilizan para asignar valores a variables.
| Operador | Ejemplo | Igual que  |
| -------- | ------- | ---------- |
| =        | x = 5   | x = 5      |
| +=       | x += 3  | x = x + 3  |
| -=       | x -= 3  | x = x - 3  |
| \*=      | x \*= 3 | x = x \* 3 |
| /=       | x /= 3  | x = x / 3  |
| %=       | x %= 3  | x = x % 3  |


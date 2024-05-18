### Ejercicios de Bucles en Go

1. **Par o Impar**: Verifica si un número es par o impar.
2. **Año Bisiesto**: Determina si un año es bisiesto o no.


## Solución

#### Ejercicio 1.

```go
package main
import "fmt"

func main() {
	num := 11
	resp := num%2 == 0

	fmt.Println("¿El número", num, "es par?", resp)
}
```
#### Ejercicio 2.

```go
package main
import "fmt"

func main() {
	anio := 2100

	resp := (anio%4 == 0 && anio%100 != 0) || anio%400 == 0
	fmt.Printf("¿Es bisiesto el año %d? %t\n", anio, resp)
}
```

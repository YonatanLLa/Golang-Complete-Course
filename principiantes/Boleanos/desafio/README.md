### Ejercicios de Bucles en Go

1. **Par o Impar**: Verifica si un número es par o impar.
2. **Año Bisiesto**: Determina si un año es bisiesto o no.
3. **Palíndromo**: Valida si una palabra es un palíndromo.
4. **Número Primo**: Comprueba si un número es primo o no.
5. **Caracteres Alfabéticos**: Verifica si una cadena contiene solo caracteres alfabéticos.
6. **Signo de un Número**: Determina si un número es positivo, negativo o cero.
7. **Tipo de Triángulo**: Comprueba si un triángulo es equilátero, isósceles o escaleno dados sus lados.
8. **Divisibilidad**: Valida si un número es divisible por otro número.
9. **Acrónimo**: Verifica si una cadena es un acrónimo (por ejemplo, "NASA").
10. **Año de Jubileo**: Determina si un año es un año de jubileo.

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

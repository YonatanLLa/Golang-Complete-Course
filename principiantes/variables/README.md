# Las variables son contenedores para almacenar valores de datos.

## En Go, hay dos formas de declarar una variable.

### 1. Utilice la palabra clave var, seguida del nombre y tipo de la variable.

```go
package main
import "fmt"
var name = "yonatan llanto"

func main() {
	var lastname = "Yonatan"

	fmt.Println(name, lastname)
}
```

### 2. Utilice el signo :=, seguido del valor de la variable.

```go
package main
import "fmt"

func main() {
	name := "yonatan"
	lastname := "llanto aquino"
	age := "24"

	fmt.Println(name, lastname, age)
}
```
### 3. Declaración de variables múltiples de Golang

```go
package main
import "fmt"

func main() {
	var name, lastname, age = "yonata", "llanto", 24

	fmt.Println(name, lastname, age)
}
```
### 4. Declaración de variable de Go en un bloque

```go
package main
import "fmt"

func main() {
	var (
		name     string
		lastname string = "Yonatan"
		age      int    = 1
	)

	fmt.Println(name, lastname, age)
}
```
## Reglas de Nombramiento de Variables en Go
#### 1. Las variables deben comenzar con una letra o un carácter de guion bajo (\_).
#### 2. Un nombre de varible no puede comenzar con un digito.
#### 3. Las variables no pueden contener espacios en blanco.
#### 4. Un nombre de variable distingenentre mayúculas y minúsculas (age, Age, AGE and \_ ) son tre variables diferentes.
#### 5. Las variables no pueden contener caracteres especiales como ```!, @, #, $, %, ^, &, *, (, ), -, +```.
#### 5. No hay limite en la longitud del nombre de la variable.
#### 6. El nombre e la variable no puede ser ninguna palabra clave de Go




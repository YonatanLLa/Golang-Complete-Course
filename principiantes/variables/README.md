# Las variables son contenedores para almacenar valores de datos.

## En Go, hay dos formas de declarar una variable.

#### 1. Utilice la palabra clave var, seguida del nombre y tipo de la variable.

```go
package main

import "fmt"

var name = "yonatan llanto"

func main() {
	var lastname = "Yonatan"
	fmt.Println(name, lastname)
}
```

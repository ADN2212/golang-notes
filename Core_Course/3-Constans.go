// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

// Constans can be shadoweb:
//const myConst int = 15

//Como crear bloques de constantes secuenciales:

const (
	x = iota//Funciona incluso sin asignar nada a las demas variables, pero esta limitado al scope
	y            //= iota
	z            //= iota
)

//Nota: se puede hacer algo como const bar = iota + 5 para empezar la secuencia desde 5 en vez de 0.

const (
	z1 = iota //This will strart by zero again
	z2
)

func main() {

	//La convencion para nombrar constantes es la misma que para la variables normales
	//const myConst int = 10
	//Por el shadowing aqui abajo aparecera 10 en vez de 15
	//fmt.Printf("%v", myConst)
	//myConst = 25 Error ...
	//No se puede crear una constante con un valor que tiene que ser determinado un tiempo de ejecucion:
	//Ej:
	//const sinOf5 float32 = math.Sin(5)//./prog.go:18:25: math.Sin(5) (value of type float64) is not constant
	//Se pueden declarar constantes con todos los datos primitivos:

	//const b string = "I Will Not Change!!"
	//const pi float32 = 3.14
	//const boOl bool = true
	//fmt.Println("Hello, 世界")

	//fmt.Printf("%v\n", b)
	//fmt.Printf("%v\n", pi)
	//fmt.Printf("%v\n", boOl)

	//Nota:si se hace una operacion entre una constante y una variable, el resultado sera una variable.

	//const a = 42
	//fmt.Printf("%v, %T\n", a, a)
	//var z int16 = 27
	//Aqui podemos ver que aunque 'a' es de tipo int y 'z' de tipo int16 el compilador de Go si puede hacer inferencia de tipos con constantes:
	//Cosa que no sucedera con variables.
	//fmt.Printf("Suma = %v, Typo = %T.\n", a + z, a + z)

	//Observece como Go asigna los valores de forma secuecial.
	 fmt.Printf("%v\n", x)
	 fmt.Printf("%v\n", y)
	 fmt.Printf("%v\n", z)
	 fmt.Printf("%v\n", z1)
	 fmt.Printf("%v\n", z2)
	//Queda pendiente investigar sobre las aplicaciones de atoi

}

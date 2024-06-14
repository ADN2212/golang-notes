// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//Bolean types:
	//var n bool = false
	//fmt.Println("Hello, 世界")
	//b := 2 == 4/2
	//En Go cuando se declara una variable sin inicializarla esta tiene el valor 0 de entrada.
	//Ej:
	//var q bool//esta es false por ser bool pero false == 0
	//var k int//esta es 0 por ser int
	//var j float32
	//fmt.Printf("%v is type %T.\n", j, j)

	//Numeric types:
	//Existen varios tipos de enteros que dependen de la cantidad de bit disponibles para representarlos
	//Estos son:
	// int8, int16, int32, int64
	//Si se nesecitan mas grades ver https://pkg.go.dev/math/big

	//Luego estan los unit que son enteros sin signo, es decir el cero y los naturales
	//Estos tienen la misma cantidad de subtipos que los int.
	//Ej:
	//var e uint32 = 10
	//fmt.Printf("%v is type %T.\n", e, e)

	//En Go hay dos tipos de flotantes: float32 y float64
	//EJ:

	//n := 13.7e72//Se puede usar notacion exponencial.
	//var pi float32 = 3.14 //Esta es la forma explicita.
	//Nota, no se pueden hacer operaciones con boleano distintos.
	//fmt.Println(n + float64(pi))//Hay que transformar de un tipo a otro.
	//fmt.Printf("%v is type %T.\n", pi, pi)

	//Por ultimo tenemos los numero complejos:
	//Ej:
	//var c complex64 = 10 + 3i
	//fmt.Printf("%v is type %T.\n", c, c)

	//Go tambien puede inferirlos:
	//c := 3.5 + 4.5i
	//Tambien se pueden crear como sigue:
	//c := complex(5, 3)
	//fmt.Printf("%v is type %T", c, c)
	//Se pueden extraer las partes reales e imaginarias de un numero complejo como sigue:
	//fmt.Printf("Parte real %v (tipo %T) parte imaginaria %v (tipo %T).\n", real(c), real(c), imag(c), imag(c))

	//Cadenas de texto:
	s := "This is a string"
	fmt.Printf("%v, %T", string(s[2]), s[2]) //El valor que aparecera sera el de la letra en ascci, por eso se use la funcion string()
	Como en otros lenguajes, las cadenas de texto son inmutables, por lo que hacer algo como s[2] = 'a' provocara un error.
	Como en otros leguajes las strs se pueden sumar.
	Otra operacion que se puede hacer es optener un array con los codigos ascci de cada caracter de la str:
	b := []byte(s)
	fmt.Printf("%v\n", b)
	
	//Runes: Que son ?
	//Los runs en Golang son types alias para los int32, there is an example:
	//run := 'r'//a single character ...
	//fmt.Printf("%v is type %T.\n", run, run)
	
	

}
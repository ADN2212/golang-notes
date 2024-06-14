package main

import "fmt"

//Aqui se pueden declarar variables, pero con la primera forma:

var pi float32 = 3.141516

//Aqui fuera se pueden declarar bloques de variables como sigue:

var (
	name   string  = "Nina"
	age    int     = 56
	heithg float32 = 1.32
)

//Ojo, estas no nesesariamente deben estar interrelacionadas.

//Nota: las variables en Go siempre deben ser usadas, de lo contrario el
//programa se rompe.

//Nota2: Si una variable es declarada con letras mayusculas es exportable hacia fuera del package
// var E float64 = 2.7182818284590452353602874713527

//Investigar sobre las combenciones para nombrar variables...

func main() {
	//Hay tres formas de declarar variables:

	//var a int
	//var b int
	//a = 10
	//b = 15

	//var a int = 20
	//var b int = 15

	//El compilador puede darce cuenta de que tipo es la variable que se declara
	//Si se usa la siguente sintaxys:
	//a := 20.
	//b := 15

	//fmt.Println(int(a) + b)//No se permiten adiciones entre tipos numericos distintos.
	//fmt.Printf("La variable 'a' tiene el valor %v y el tipo %T", a, a)

	//no se puede redeclarar una variable en el mismo scope:
	//a := 32 //error => no new variables on left side of :=

	//Si declaro una variable con el mismo nombre que una externa
	//se aceder a esta antes que a la mas externa, esto es Shadowing:
	//name := "Juan"
	//fmt.Printf("La %v tiene %v años y medie %v m de altura.", name, age, heithg)

	//Se puede cambiar de un tipo a otro como sigue:
	//var z int = 2
	//var j float32
	//j = float32(z)
	//fmt.Printf("La variable j tiene el valor %v y el tipo %T.\n", j, j)

	//Que pasa al comvertir de str a int?
	 var arrobaInt int = 42
	 var arrobaStr string
	 arrobaStr = string(arrobaInt)
	 arrobaStr = fmt.Sprint(arrobaInt)
	 fmt.Println(arrobaStr)
	 fmt.Printf("%v", arrobaStr)
	//Ver el strcov package.

}
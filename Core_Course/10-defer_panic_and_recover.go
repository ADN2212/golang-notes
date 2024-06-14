package main

import (
	"fmt"
)

func main() {

	//DEFER:
	//fmt.Println("Start")
	//defer fmt.Println("Middle") //Retraza la ejecucion de esta linea hasta el final de la funcion.
	//defer fmt.Println("End")//Si hay mas de una se comporta como una stack.

	//Que sucedera cuando se ejecuten estas lineas:
	//a := "Pirmer valor"
	//defer fmt.Println(a) //Se mantiene el valor que tenia en el momento.
	//a = "Segundo valor"
	//fmt.Println(a)

	//PANIC:
	//En Go no hay excepciones, se puede usar la palabra panic para indicar que algo que afectara la ejecucion del codigo ha ocurrido:
	fmt.Println("Hola")
	defer fmt.Println("What?")//La funcion ejecutara los defers antes del panic.
	panic(" Bobo en la vereda :( ") //Esto detiene la ejecucion del programa.
	fmt.Println("Adios...")         //Ne se ejecuta
	//Se puede usar la funcion recover() para salvar un programa de detenerce definitivamente por un panic.
}
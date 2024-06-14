package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World")

	//Interfaces:
	//myCounter := IntCounter(0)       //Inicializacion del contador
	//var inc mainCounter = &myCounter //Uso de la interfas para poder usar el metodo.

	//for i := 0; i < 10; i++ {
	//fmt.Println(inc.Increment())
	//	inc.Increment()
	//	fmt.Printf("Incrementing ... + 2 = %v\n", myCounter)
	//	if i%3 == 0 {
	//		inc.Decrement()
	//		fmt.Printf("Decrementing... -1 = %v\n", myCounter)
	//	}
	//}
	//fmt.Println(myCounter) //Vease como el valor del counter cambio

	//The empty interface es una interface sin metodos
	//esto le permite hacer casting de cualquier tipo a esta.
	//Imaginemos que hemos resivido el valor de n y no sabemos de que tipo es,
	//la interfaz vacia no sera util para almacenarlo antes de saber su tipo
	//para luego comvertirlo, Ej:

	var bar interface{} = "What am i ?"
     
	switch bar.(type) {
	case int:
		fmt.Println("Is an integer")
	case string:
		fmt.Println("Is a string")
	case float64:
		fmt.Println("Is a real number")
	default:
		fmt.Println("What is that ???")
	}

}

// Interfaces:
// Asi como las structs sirven para agrupar variables relacionadas entre si,
// las interfaces sirven para agrupar metodos que estan relacionados entre si.
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	(*ic) += 2
	return int(*ic)
}

//Tambien se pueden crear interface que estan compuestas por otras interfaces para agrupar mas comportamientos, Ej:

type Decrementer interface {
	Decrement() int
}

func (ic *IntCounter) Decrement() int {
	(*ic) -= 1
	return int(*ic)
}

// Esta es una interfaz de iterfaces:
type mainCounter interface {
	Incrementer
	Decrementer
}

//Pendiente repasar este tema.
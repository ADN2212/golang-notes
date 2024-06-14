package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World")

	//Loops:
	//En Go solo existe un tipo de loop, el for:
	//Notece como no es nesesario el uso de parentesis
	//for j := 0; j < 10; j = j + 1 {
	//	fmt.Println(j + 1)
	//}

	//Se pueden inicializar mas de una variable en el for:
	// for a, b := 0, 3; a + b < 35; a, b = a + 1, b + 1 {
	// 	fmt.Println(a, b)//el scope de a y b es el cuerpo del loop
	// }

	//No es obligatorio que haya una inicializacion en el for:

	//z := 0//el scope de z es la main func
	//Hay que dejar el punto y como para que se sepa que se esta dejando fuera la inizializacion:
	// for ; z < 25; z = z + 5 {
	// 	fmt.Println(z)
	// }

	//El siguiente for se parece mas a un while:
	//var q int = 100
	//Observece como aqui no se nesecita el ;
	// for q > 0 {
	// 	fmt.Println(q)
	// 	q = q - 1
	// }

	//Se puede hacer un ciclo de iteraciones identerminadas como sigue:

	//r := 0
	//Este es equivalente a un 'While True' en python.
	// for {
	// 	r++
	// 	if r % 2 == 0 {
	// 		continue//Sata a la proxima iteracion.
	// 	} else {
	// 		fmt.Println(r)
	// 	}
	// 	if r > 10 {
	// 		break //detiene el ciclo como en python
	// 	}
	// }

	//tambien es posible hacer loops anidados:
	// Loop:
	// for i := 1; i <= 3; i++ {
	// 	for j := 1; j <= 3; j++ {
	// 		fmt.Printf("%v X %v = %v\n", i, j, i*j)
	// 		if i*j >= 3 {
	// 			break Loop //Why is this posible?
	// 		}
	// 	}
	// }
	//Working whit collections:
	//s := []int{1, 2, 3}
	//fmt.Printf("%T", s)
	//el range de Go permite acceder al indice y el valor en cada iteracion:
	// for i, v := range s {
	// 	fmt.Printf("index = %v, value = %v\n", i, v)
	// }

	//Tambien funcionara con maps:

	populations := map[string]int{
		"RD":    1000000,
		"Haiti": 900000,
		"Esp":   5000000,
		"USA":   150000000,
	}

	//Si solo se quiere el valor podemos dejar el primer valor como _
	for name, population := range populations {
		fmt.Printf("%v has %v\n", name, population)
		//fmt.Println(population)
	}

}

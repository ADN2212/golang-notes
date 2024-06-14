package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World")

	//Maps:
	//Los maps son como los diccionarios de Python o los JSONs
	//Se pueden usar casi todos los tipos de datos como llaves menos los slices y ...
	//team := map[string]int{
	//	"ADN":    28,
	//	"Ksoto":  25,
	//	"Anfer":  19,
	//	"Edwin":  22,
	//	"Samael": 24,
	//}
	//Estos tambien pueden ser creados usando la funcion make()

	family := make(map[string]string)
	//var family map[string]string
	family = map[string]string{"Nina": "Curz Maria", "hood": "Ericson Gabriel"}

	//family["Nina"] = "Cruz Maria"

	//Se pueden agregar nuevos valores:
	//family["YO"] = "Adonis Nunez"
	//fmt.Println(family)
	//Ojo: los maps no tienen un orden definido
	//Se pueden eliminar elementos del map usando la funcion delete:
	//delete(family, "YO")
	//fmt.Println(family)

	//Si una llave no existe dentro de un map al buscarla se optendra el valor por default, EJ:
	//fmt.Println(family["Manetha"]) //Esto retorna una str vacia por ser de tipo string
	//Para evitar confucion a la hora de saber si un valor esta o no en un map hacemos:
	//manetha, ok := family["Manetha"]//ok es una convencion
	//fmt.Println(manetha, ok)//aqui ok = false si el valor no esta.
	//Tambien es posible usar la funcion len con un map
	//fmt.Printf("Your family has %v members", len(family))

	//Otra cosa a tomar en cuenta es que las copias de un map apuntan al mismo espacio en memoria, Ej:
	secondFamily := family
	secondFamily["Pa"] = "Juan Nunez"
    //Observece como "Pa" fue agregado al map family
	fmt.Println(family)

}

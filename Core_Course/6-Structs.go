package main

import (
	"fmt"
	"reflect"
)

// Las structuras son grupos de datos relacionados entre si
// algo parecido a los objetos, pero ojo en Go no hay POO
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

//Las mayusculas permitiran que sea exportable

func main() {
	//aDoctor := Doctor{
	//	number:     10,
	//	actorName:  "Maria Elena",
	//	companions: []string{"Henoch", "YO", "Nikauris"},
	//}

	//Tambien podemos declararla de la siguiente manera:
	//Aqui se estan obviando los nombres de los campos, Go sabra cual campo es por la posicion en la que se encuentra.
	//El autor del curso recomienda no usarla, dado que puede generar errores al cambiar la structura.
	// aDoctor := Doctor{
	// 	10,
	// 	"Maria Elena",
	// 	[]string{"Henoch", "YO", "Nikauris"},
	// }

	//Como en POO podemos usar la notacion del punto:
	// fmt.Println(aDoctor.Number)
	// fmt.Println(aDoctor.ActorName)
	// fmt.Println(aDoctor.Companions[2])

	//Se pueden declarar variables con structuras anonimas como sigue:
	//anotherDoctor := struct{ name string }{name: "Tontin"}
	//fmt.Printf("%v is a %T", anotherDoctor, anotherDoctor)

	//A diferencia de los maps las copias de las structs no apuntan a el mismo espacio en memoria
	//por lo que al hacer:
	//lastDoc := anotherDoctor
	//lastDoc.name = "Juan"
	//fmt.Println(anotherDoctor.name)//Vease como no cambia el valor.
	//Ojo: Se puede hacer que si apunten al mismo espacio de memoria usando pointers.

	//Golang no soporta POO por lo que no hay herencia, pero se pueden conseguir comportamientos parecidos usando composition, EJ:

	type Animal struct {
		Name   string `This is a tag and adds info to the struct`
		Origin string
	}

	type Bird struct {
		Animal //Aqui se estan 'heredando' las caracteriticas del type animal
		Speed  float32
		CanFly bool
	}

	//Ahora podemos crear una instancia de la structura bird
	// b := Bird{}
	// b.Name = "Avestruz"
	// b.Origin = "Africa"
	// b.Speed = 75.32
	// b.CanFly = false

	//Pero lo que realmente esta pasando es algo como esto:
	// b := Bird{
	// 	Animal: Animal{
	// 		Name:   "Agilla",
	// 		Origin: "Ei Cibao",
	// 	},
	// 	Speed:  150.55,
	// 	CanFly: true,
	// }
	
// 	fmt.Println(b)

// 	//How to get the tag info:
// 	 t := reflect.TypeOf(Animal{})
// 	 field, otherThing := t.FieldByName("Name")
// 	fmt.Println(field, otherThing)

	//Como iterar a travez de una struct:

	a := Animal{"Perro", "Canis"}
	values := reflect.ValueOf(a)//Para el caso, los valores de la strcut
	bar := values.Type()//El tipo de donde vienen los valores
	fmt.Println(values.NumField())//la cantidad de campos
	fmt.Println(bar.Field(1))//el campo por el indice.
 	fmt.Println(a[bar.Field(1).Name])



}

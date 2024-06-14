package main

import (
	"fmt"
)

func main() {

	//Pointers:
	//a := 50
	//b := a
	//fmt.Println(a, b)
	//a = 25
	//Aqui podemos ver que cambiar a no implica cambair b, porque estos son 'value types' y apuntan a lugares diferentes de la memoria.
	//fmt.Println(a, b)

	//var a int = 50
	//var b *int = &a // *int significa que b es el puntero de un entero y &a es la direcion en memoria de a
	//fmt.Printf("El valor de a es %v y está ubicado en %v\n", a, b)
	//a = 25
	//Cuando se hace *b se optiene el valor que esta guardado en esta direccion, esto se llama 'dereference'
	//fmt.Println(a, *b)//Son iguales porque apuntan al mismo valor.

	//En Go no esta permitido hacer pointer aritmetic, pero en coso de que se desee ver el 'unsafe package'

	//Pointers y arrays:
	//array := [3]int{1, 2, 3}
	//b := &array[0] //Puntero al primer valor del array
	//c := &array[1] //Puntero al segundo valor del array
	//fmt.Printf("%v, en %p esta %v y en %p esta %v\n", array, b, *b, c, *c)

	//Pointers y structs
	//var ms *myStruct //ms es un puntero que apunta a un dato de tipo myStruct
	//fmt.Println(ms)  //Que es ms antes de asignarle un valor ???
	//Si en este punto se hiciera (*ms).foo lanzaria una exepcion.
	//ms = &myStruct{foo: 42} //Esta es la direccion en memoria de un dato de tipo ....,
	//esto tambien se puede hacer con 'new(myStruct)' pero no se puede inicializar en la misma linea.
	//fmt.Println(ms)        //Porque no aparece la forma numerica de la direccion??
	//fmt.Println((*ms).foo) //aqui, se le hace dereference a ms y luego se accede a el valor foo de la struct.
	//(*ms).foo = 55
	//Sin embargo Go permite hacer lo siguiente:
	//ms.foo = 60//Esto es 'syntaxys sugar', pero la anterior es lo que realmente ocurre.
	//fmt.Println(ms.foo)

	//Pointers and Slices:
	a := []int{5, 6, 7}
	b := a
	//fmt.Println(a, b)
	a[1] = 78
	//fmt.Println(a, b)
	//Aqui vemos que al cambiar un elemento de a se cambia tambien un elemento en b, esto es porque
	//b esta apuntando al mismo espacio en memoria de a, es decir que &a = &b
	//por lo que &a[i] = &b[i] para todo i en [0, len(a ó b)]
	//lo mismo pasa con los valores, *&a[i] = *&b[i] = a[i] = b[i], Ej:
	fmt.Println(*&a[1], *&b[1])

	//Esto tambien sucede con los maps:
	// a := map[string]string{"Foo": "foo?", "Bar": "bar?"}
	// b := a
	// fmt.Println(a, b)
	// a["Bar"] = "Yes, bar"
	// fmt.Println(a, b)

}

type myStruct struct {
	foo int
}

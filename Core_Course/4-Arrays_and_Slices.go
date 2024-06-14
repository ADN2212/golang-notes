package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World")
	//Arrays:
	//Hay que espesificar la longitud y el tipo de dato de cada elemento:
	//Los elementos de un array son contiguos en la memoria
	//grades := [3]int{80, 75, 20}
	//grades := [...]int{80, 75, 65} //Al poner los tres puntos el array es declarado con longitud igual a su numero de elementos.
	//fmt.Printf("Calificaciones => %v", grades)

	//Se pueden declarar arrays sin necesidad de asignar valores a sus indices:
	//var students [3]string
	//Se pueden asignar valores con la sintaxys basica:
	//students[0] = "Juan"
	//students[1] = "Nina"
	//students[1] = 22//Error "/prog.go:17:16: cannot use 22 (untyped int constant) as string value in assignment"
	//fmt.Printf("Thesse are the students => %v we have space for %v students.\n", students, len(students))

	//Tambien es posible tener arrays dentro de arrays para crear matrices:
	//Si no se asignan valores estos seran inicializados en funcion de su tipo, Ej para int o float es 0.
	//var matriz [3][3]int
	//Obviamente se pueden asignar valores haciendo doble sub indice.
	// matriz[0][0] = 3
	// matriz[0][1] = 4
	// matriz[0][2] = 5
	// matriz[1][0] = 78
	// matriz[1][1] = 25
	// matriz[1][2] = 33
	// matriz[2][0] = 28
	// matriz[2][1] = 47
	// matriz[2][2] = 32
	// fmt.Printf("%v", matriz)

	//En Go los arrays no son referencia si no valores (en contraposicion a JS por ejemplo)
	//Por lo que si creamos una copia de un array, esta no apuntará al mismo espacio de memoria que el original

	//var matriz1 [3][3]int
	//matriz2 := matriz1
	//matriz2[0][0] = 10000
	//Notece como al cambiar el primer elemento de la segunda matriz no se cambia el de la primera.
	//Nota: si quisieramos que matriz2 sea una shallow copy de matriz1 lo podriamos lograr usando pointers (se vera mas adelante en el curso)
	//fmt.Printf("%v\n", matriz1)
	//fmt.Printf("%v\n", matriz2)

	//Slices:
	//Estos son como los de Python

	//a := []int{1, 2, 3, 4}
	//Los slices si apuntan al mismo espacio en memoria que su array padre
	//b := a[2:]
	//b[0] = 10 //Esto cambia el primer elemento de a y por tanto el primero de c
	//c := a[1:3]
	//fmt.Printf("%v, %v, %v\n", b, c, cap(a))
	//Queda pendiente varificar la diferencia entre len() y cap()

	//Tambien se pueden crear slices usando la funcion make()
	//Aqui, el primer argumento es el array del que se hace el slice, el segundo la longitud y el tercero la capacidad
	//s := make([]int, 3, 100)
	//Vease como s tiene inicialmente solo tres elementos (len), pero puede ser ampliado hasta 100 (cap).
	//fmt.Printf("Original slice => %v, len = %v\n", s, len(s))
	//Para agregar mas elementos a s usamos la funcion append()
	//s = append(s, 4, 5, 200)//Esto deja clara la diferencia entre cap y len, ver pag 69 del libro Go in action
	//fmt.Printf("New Slice = %v, Len = %v y Cap = %v\n", s, len(s), cap(s))

	//Mas sobre Slices y Arrays:

	//Sea 'a' un array de longitud 'len' y 's[i:j]' un slice de este con 0 <= i <= j
	//la longitud 'l' de s es:
	//l = j - i
	//Y la capacidad 'c' es:
	//c = len - i

	//Ejemplo:

	// a := [10]int{}
	// fmt.Println(a) 
	// s := a[3:5]//l = 5 - 3 = 2
	// fmt.Println(cap(s))//c = 10 - 3 = 7

	//Al usar append con un slice que esta a su maxima capacidad se duplicara la capacidad del array de fondo 
	//Ej:

	s := []int{2, 3}
	fmt.Println(cap(s))
	s = append(s, 3)
	fmt.Println(cap(s))
	s = append(s, 7, 8)
	fmt.Println(cap(s))

	//Esto sucede para cap < 1000, de ahi en adelante crece a razon de un 25%

}

package main

import "fmt"

func main() {
	//fmt.Println("Hello World")
	//Functions:
	//saySomething("HolaaaaaaaaaAAaaaa")
	//a := 5
	//product(&a, 3)
	//fmt.Println(a) //has changed
	//sumAll(4, 3, 5, 6, 7)
	//resDir := sumAll2(4, 3, 5, 6, 7)
	//fmt.Printf("En %p esta el valor %v", resDir, *resDir)
	//dif := sustrac(10, 5)
	//fmt.Println(dif)
	//cotient, error := divide(10, 2)//Unpacking variables...
	//fmt.Println(cotient, error)

	//En Golang tambien hay IIFE:
	//y funciones anonimas:
	//func() {
	//	fmt.Println("Im here but Im not")
	//}()

	//Y se pueden asignar funciones a variables:
	//Ojo al tipo, se deben espesificar los tipos de los parametros y del valor que retorna.
	//var foo func(int, int) int = func(a, b int) int {
	//	fmt.Println("Im foo!!")
	//	return a + b
	//}
	//foo(3, 4)

	yo := person{
		name: "Adonis",
		age:  28,
	}

	yo.present() //Like a proto POO
	yo.bridthDay()
	fmt.Println(yo.age)//Notece como el valor de yo.age cambio.
}

// Una funcion puede tener o no parametros, en caso de, hay que decir el tipo de estos:
func saySomething(msg string) {
	fmt.Println(msg)

}

// Si varios de los parametros son del mismo tipo, se puede hacer lo siguiente:
// func product(a, b int) {
func product(a *int, b int) {
	fmt.Println(*a * b)
	*a = 10 //Si a la funcion se le pasa el puntero de a, sera posible mutar su valor desde dentro de la funcion.
	//con los valores normales como b esto no pasa.
	//Pasar pointers en vez de valores normales se traduce en mejor performance, ya que no se hacen copias de los valores.
	//Toda funcion que use pointers is not a pure function.
}

// Variatics parameters:
// Como en JS o Python una funcion puede tener un numero ideterminado de parametros:
func sumAll(nums ...int) { //Solo puede haber un variatic parameter y debe estar al final.
	//fmt.Printf("%v is type %T", nums, nums)//Un slice
	sum := 0
	for _, n := range nums {
		sum += n
	}

	fmt.Println(sum)
}

// Si la funcion es no vacia, debe espesificarce el tipo de valor que retorna.
func sumAll2(nums ...int) *int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return &sum //Tambien es posible retornar punteros. Que implicaciones tiene esto?
	//Al hacer esto Go no liberara el espacio en memoria despues de finalizar la ejecucion de la función.
}

// Tambien se pueden nombrar el valor que retorna una funcion en la declaracion de esta:
func sustrac(x, y int) (rest int) {
	//fmt.Println(rest)
	rest = x - y
	return //Go entiende que rest es lo que tiene que retornar.
	//Esta sintaxys no se usa mucho
}

// Tambien una funcion puede retornar mas de un tipo de valor.
func divide(a, b int) (float64, error) {
	if b == 0 {
		//panic("Zero division error")
		return 0.00, fmt.Errorf("Zero division error")
	}

	return float64(a / b), nil
}

// Methods:
// Los metodos son funciones que se le aplican a structuras, lo mismo que los metodos en POO:
type person struct {
	name string
	age  uint
}

// Aqui p es un value type, no una referencia.
// En caso de que se requiera, usar pointers.
func (p person) present() {
	fmt.Printf("Hola, soy %v y tengo %v años.\n", p.name, p.age)
}

func (dirToP *person) bridthDay() {
	(*dirToP).age += 1//Para el compilador (*dirTop) = p, para simplificar la escritura del codigo.
	fmt.Printf("Happy birth day %v, you are %v year old!!\n", (*dirToP).name, (*dirToP).age)
}

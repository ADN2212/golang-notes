package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("Hello World")
	//If statements:
	//Un if basico
	//if true {
	//	fmt.Println("Hello World")
	//}
	//Uno un poco mas complejo:
	
	// friends := map[int]string{
	// 	1: "JDK",
	// 	2: "Hito",
	// 	3: "Gian",
	// }

	//Lo que sucede aqui es que despues de declararce 'ok' esta se usa como coindicion.
	if w, ok := friends[35]; ok {
		fmt.Printf("%v Is your friend", w)
	} else {
		fmt.Println("You dont have this friend")
	}

	//An example whit a guessing game:
	//number := 100
	//guess := 150
	//Observecce como en caso de que guess sea menor que number no se ejecuta la funcion alwaysTrue
	//Esto es porque en un or vasta con que un solo valor sea true, luego Golang se ahora la ejecucion de la funcion para mejorar el performance.
	//Investigar sobre short circuiting:
	//Esto tambien funciona con el && ya que si uno de los valores es falso entonces toda la expresion lo será
	// if guess < number || alwaysTrue() {
	// 	fmt.Println(" Ma' pa'lante' ")
	// }
	// if guess > number {
	// 	fmt.Println(" Ma pa' tra' ")
	// }
	// if guess == number {
	// 	fmt.Println(" Ese mimo' e ")
	// }
	//By the way, los operadores de comparacion son solo 6 como en la mayoria de lenguajes (>, <, >=, <=, ==, !=) y solo funcionan con numeric tyepes.
	// Logicos: && = and, || = or ! = not
	//Tambien existen los else statements:
	//if !alwaysTrue() {
	//	fmt.Println("Do something")
	//} else {
	//		fmt.Println("Do anotherthing ..")
	//}
	//Y las condiciones intermedias: if {} else if {} else {}

	//A la hora de usar flotantes hay que tener en cuenta que Go (como todos los lenguajes) usa aproximaciones, Ej:
	myNum := 0.11321321321321321
	mySamenNum := math.Pow(math.Sqrt(myNum), 2)
	//fmt.Printf(" %v = %v ?", myNum, mySamenNum)//Vease como para valores muy extensos deja de ser igual
	// if myNum == mySamenNum {
	// 	fmt.Println("they are equal")
	// } else {
	// 	fmt.Println("They are not the same")
	// }

	//Por lo deberiamos truncar (redondear) numeros muy extensos para verificar su igualdad.
	if math.Abs(myNum-mySamenNum) < 0.001 {
		fmt.Println("they are almost the same")
	} else {
		fmt.Println("they are enough diferent to not be the same")
	}

}

func alwaysTrue() bool {
	fmt.Println("I Will always return true")
	return true
}
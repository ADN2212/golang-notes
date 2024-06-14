package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World")

	//Swicth statements:
	//Como los swicths the JS

	//SexualOrientations := map[int]string{
	//	1: "Hetero",
	//	2: "Bi",
	//	3: "Homo",
	//}

	//yours := SexualOrientations[1]

	//Tambien se puede hacer la condicion despues de la asignacion:
	//switch yours := SexualOrientations[3]; yours {
	//Pueden englobarce varios valores en un solo caso
	//case "Bisexual", "Homo":
	//	fmt.Println("God hates you!")
	//case "Hetero":
	//	fmt.Println("You are ok")
	//default:
	//	fmt.Println("What are you?")
	//}

	//Tambien exista la tagless sintax;
	//switch {
	//case yours == "Homo" || yours == "Bi":
	//	fmt.Println("God hates you!")
	//   //fallthrough
	//case yours == "Hetero":
	//	fmt.Println("You are ok")
	//default:
	//	fmt.Println("What are you?")
	//}
	//Observece como no se usa el break statement dado que está implicito,
	//pero si no se quiere salir del swictch despues de que se cumpla un caso solo hay que poner fallthrough.

	//Tambien se pueden usar para decidir que hacer en caso del tipo de la variable:

	var n interface{} = 10

	switch n.(type) {
	case int:
		fmt.Printf("%v is an integer\n", n)
		break
		fmt.Println("Maybe you dont want to see this info")
        //El break implicito esta aqui.
	case float64:
		fmt.Printf("%v is a float", n)

	case string:
		fmt.Printf("%v is a string", n)

	default:
		fmt.Println("I dont know waht it is")
	}

}

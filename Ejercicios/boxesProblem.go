package main

import (
	"fmt"
	"sort"
)


func main(){
	//fmt.Println("Hellow")
	//Data:
	// box1 := box{large: 1, width: 1, high: 1,}
	// box1.CalcVolumen()//Puedo hacer esto al declarar la box?
	// box2 := box{large: 2, width: 2, high:2 ,}
	// box2.CalcVolumen()
	// box3 := box{large: 3, width: 3, high: 3,}
	// box3.CalcVolumen()
	// boxes1 := []box{box3, box1, box2}

	//Dado que los metodos Len, Less y Swap hace referencia al tipo boxList, hay que transformar el array de boxes a este tipo:
	//sort.Sort(boxList(boxes1))

	// b1 := box{1, 1, 1,0,}
	// b2 := box{2, 2, 2,0,}
	// b3 := box{3, 1, 3,0,}

	boxes2 := boxList(
		[]box{
			box{1, 1, 1,},
			box{2, 2, 2,},
			box{3, 1, 3,},
	})
	
	sort.Sort(boxes2)
	//fmt.Println(boxes2)

	// b1 := box{1, 1, 1,}
	// //b1.CalcVolumen()
	// b2 := box{2, 2, 2,}
	// //b2.CalcVolumen()
	// b3 := box{2, 10, 2,}
	// //b3.CalcVolumen()

	//boxes3 := boxList([]box{b1, b2, b3})

	//sort.Sort(boxes3)

	//Finalmente interar por el slice ordenado y ver si todas las cajas caben una dentro de otra:

	if !finxInOneBox(boxes2) {
		fmt.Println("They dont fix")

	} else {
		fmt.Println("They fix!!")
	}

}


//Solucion del problema de las cajas con Golang:
//ver 

//Esta es la estructura que representa las cajas:
type box struct {
	large int
	width int
	high int
	//volume float64
	//Aunque el algun contexto pueda llegar a ser necesario que el volumen sea una 'propiedad' de la caja, para este caso no lo es
}

//Metodos o Funciones:
//Esto es sintaxys sugar ya que si no estuviera el puntero no se cabiaria el valor real.
func (boxPointer *box ) CalcVolumen() float64 {
	//Este metodo calcula el volumen de la caja y asigna el valor al valor de la caja en la cual se ejecuta.
	//(*boxPointer).volume = float64((*boxPointer).large * (*boxPointer).width * (*boxPointer).high)
	return float64((*boxPointer).large * (*boxPointer).width * (*boxPointer).high)
}

func fixInto(box1Pointer *box, box2Pointer *box) bool{
	//Esta es una funcion que comprueba si la primera caja cabe en la segunda.
	return 	(*box1Pointer).large < (*box2Pointer).large && (*box1Pointer).width < (*box2Pointer).width && (*box1Pointer).high < (*box2Pointer).high
}


//Para ordenar la lista de empleados:

type boxList []box//Un tipo que sirve como abstraccion de slice de cajas.

//Recordar que los Slices son valores referenciales, es decir que si una de estas funciones (metodos) afecta algun valor de su argumento, tambien afectara el original.


//Len y Less y Swap son metodos que hay que definir para que se puedan ordenar los slices.
func (bl boxList) Len() int {
	return len(bl)
}

//Este metodo determina el orden del slice final
func (bl boxList) Less(i, j int) bool {
	return bl[i].CalcVolumen() < bl[j].CalcVolumen()
}

//Este determina como intercambiar en caso de que un objeto sea menor que otro:
func (bl boxList) Swap(i, j int) {
	 bl[i], bl[j] = bl[j], bl[i]	
}


func finxInOneBox(bl boxList) bool{

	len := len(bl)

	for i := 0; i <= len - 2; i++{
		//fmt.Println(boxes1[i])
		//fmt.Println(boxes1[i + 1])
		if !fixInto(&bl[i], &bl[i + 1]){
			return false
			break
		}
	}

	//fmt.Println("They Fix")
	return true

}

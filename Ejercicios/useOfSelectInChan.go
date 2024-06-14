//Ejemplo de como prosesar info de varios canales:
package main


import "fmt"
import "sync"


func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)
	ge := sync.WaitGroup{}
	ge.Add(3)
	go generateEvens(10, evenChan, &ge)
	go generateOdds(10, oddChan, &ge) 
	go getNumbers(evenChan, oddChan, &ge)
	ge.Wait()

}

func getNumbers(evenChan <- chan int, oddChan <- chan int, geDir *sync.WaitGroup){
	// for even := range evenChan{
	// 	fmt.Println(even)
	// }
	//Usando un for sin condicion (as while true) podemos optener la info de los dos canales como sigue:
	
	//Declare estas variables aqui para poder detener el ciclo
	//En el video de referencia se uso un canal para detenerlo.
	//ver 'https://www.youtube.com/watch?v=NFe33QUd6j8'
	var even int
	var odd int
	thereAreEvens := true
	thereAreOdds := true

	for {
		//Este es un select especial para iterar sobre los valores de un canal:
		select {	
			case even, thereAreEvens = <- evenChan:
				if thereAreEvens {
					fmt.Printf("%v es par.\n", even)
			}
			case odd, thereAreOdds = <- oddChan:
				if thereAreOdds {
					fmt.Printf("%v es impar.\n", odd)
			}
		}  
		if !thereAreEvens && !thereAreOdds {
			break
		}
		//fmt.Println(thereAreEvens, thereAreOdds)
	}

	//Ojo, si el ciclo de arriba no tuviera una 'break condition' el programa no se detendria, ya que esta gorutine no tendria fin.
	(*geDir).Done()

}

func generateEvens(max int, evenChan chan <- int, geDir *sync.WaitGroup){
	n := 2
	for n <= max {
		evenChan <- n
		n += 2
	}
	close(evenChan)
	(*geDir).Done()
}

func generateOdds(max int, oddChan chan <- int, geDir *sync.WaitGroup){
	n := 1
	for n <= max {
		oddChan <- n
		n += 2
	}
	close(oddChan)
	(*geDir).Done()
}

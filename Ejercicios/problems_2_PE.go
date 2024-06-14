package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){

	//Solucion secuencial
	// inicio := time.Now()	
	// prevFib := 1
	// actualFib := 2
	// //numFibElems := 2
	// evenSum := 2

	// for actualFib < 4000000 {
	// 	actualFib, prevFib = actualFib + prevFib, actualFib		
	// 	//numFibElems += 1
	// 	//fmt.Println(prevFib, actualFib)
	// 	if actualFib % 2 == 0 {
	// 		fmt.Printf("%v is an even fib num.\n", actualFib)
	// 		evenSum += actualFib
	// 	}
	// }

	// fmt.Println(time.Since(inicio))
	// fmt.Printf("La ejecución termino en %v.\n", time.Since(inicio).Microseconds())
	
	//Solucion paralela o concurrente?
	inicio := time.Now()
	fibChan :=  make(chan int)
	grupoEspera := sync.WaitGroup{}
	evenSum := 0

	//Creamos una gorutine que envie los fibNums:
	grupoEspera.Add(2)
	go func(ch chan <- int) {	
		prevFib := 1
		actualFib := 2
		fmt.Println(actualFib)
		for actualFib < 100 {
			actualFib, prevFib = actualFib + prevFib, actualFib		
			fmt.Println(actualFib)
			ch <- actualFib
		}
		close(ch)
		grupoEspera.Done()
	}(fibChan)
	//Y una que sume los que son pares:
	go func (ch <- chan int){
		//Esta gorutine se encarga de sumar los valores que envia la otra.
		for fib := range ch {
			if fib % 2 == 0 {
				fmt.Printf("%v is an even fib num.\n", fib)
				evenSum += fib
			}
		}
		grupoEspera.Done()
	}(fibChan)
	grupoEspera.Wait()
	fmt.Println(time.Since(inicio))
	fmt.Printf("La suma es %v\n", evenSum)

}

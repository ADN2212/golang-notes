package main


import "fmt"
import "sync"
import "time"

func collaztSecLongSec(n int) int{
	//fmt.Printf("Ejecutando...")
	long := 1
	//val := n
	for n != 1 {
		if n % 2 == 0{
			n = n / 2
		} else {
			n = 3*n + 1
		}
		long += 1
	}
	//fmt.Printf("Gorutine terminada para la secuencia de Collazt de %v.\n", val)
	//fmt.Println("Done for ", val)
	return long
}

func collaztSecLongSecConcurrent(n int, numDir *int, longDir *int, geDir *sync.WaitGroup){
	currentLong := 1
	val := n
	for n != 1 {
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		currentLong += 1
	}

	if currentLong > *longDir {
		*numDir = val
		*longDir = currentLong
	}

	(*geDir).Done()//Damos la Gorutine por terminada.
	fmt.Printf("Gorutine terminada para la secuencia de Collazt de %v.\n", val)
	//fmt.Println("Done for", val)
	//return long
}

func main(){
	//El proble 14 de PE sirve como ejemplo de concurrencia, ya que las longitudes de las secuencias varian
	
	//Solucion concurrente - 1:
	ge := sync.WaitGroup{} 
	var longestNum int
	var longestCollaztSecLong int
	n := 1
	start := time.Now()
	//ge.Add(1000000)
	for n < 10 {
		ge.Add(1)
		go collaztSecLongSecConcurrent(n, &longestNum, &longestCollaztSecLong, &ge)
		n += 1
	}
	
	//fmt.Println(n)
	ge.Wait()
	fmt.Println(time.Since(start))
	fmt.Printf("%v produce la mayor secuencia, esta es de %v números.\n", longestNum, longestCollaztSecLong)

	//Solucion secuencial:
	
	// longestCollaztSecLong := 0
	// var longestNum int
	// var currentLong int
	// n := 1
	// //fmt.Println(collaztSecLong(13))
	// start := time.Now()
	// for n < 1000000 {
	// 	currentLong = collaztSecLongSec(n)
	// 	if currentLong > longestCollaztSecLong {
	// 		//fmt.Println(currentLong)
	// 		longestCollaztSecLong = currentLong
	// 		longestNum = n 
	// 	}
	// 	//fmt.Println(n)
	// 	n += 1
	// }
	// fmt.Println(time.Since(start))
	// fmt.Printf("%v produce la mayor secuencia, esta es de %v.\n", longestNum, longestCollaztSecLong)

}


//Nota:Al parecer, para este caso la solucion secuencial supera la concurrente...



//IIFE:
		// go func (){
		// 	currentLong := 1
		// 	num := n
		// 	for num != 1 {
		// 		if num % 2 == 0 {
		// 			num = num / 2
		// 		} else {
		// 			num = 3*num + 1
		// 		}
		// 		currentLong += 1
		// 	}
		// 	if currentLong > longestCollaztSecLong {
		// 		longestNum = n
		// 		longestCollaztSecLong = currentLong
		// 	}
		// 	fmt.Println("Done for", n)
		// 	ge.Done()
		// }()


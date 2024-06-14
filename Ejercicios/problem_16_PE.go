package main

import (
	"fmt"
	"sync"
	//"time"
	//"math"
	"math/big"
	"strconv"
	"time"
)

func main() {

	//Ver 'https://pkg.go.dev/math/big'
	// //Esto es un puntero no un valor es un puntero:
	// start := time.Now()
	// towToThe1000 := big.NewInt(0)
	// towToThe1000 = towToThe1000.Exp(big.NewInt(2), big.NewInt(1000), nil)
	// //fmt.Printf("%v, %T.\n", towToThe1000, towToThe1000)//Ver el tipo
	// strValue := towToThe1000.String()
	// //fmt.Printf("%v, %T.\n", strValue, strValue)
	//  sum := 0 
	//  current := 0 
	//   for _, d := range strValue {
	//   	current, _ = strconv.Atoi(string(d))
	//   	sum += current
	// }
	// fmt.Println(time.Since(start))
	// fmt.Println(sum)


	//Solucion Concurrente:
	start := time.Now()
	sum := 0
	digitChan := make(chan rune)
	ge := sync.WaitGroup{}
	ge.Add(2)
	//Una gorutine que envia los digitos:
	go func(dch chan <- rune){
		//towToThe1000 := 
		strValue := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil).String()
		for _, d := range strValue {
			 dch <- d//Enviamos el digito actual al canal
		}
		close(dch)		
		ge.Done()
	}(digitChan)
	//Una gorutine que toma los digitos y los suma:
	go func(dch <- chan rune){
		current := 0
		for strDigit := range dch {
			current, _ = strconv.Atoi(string(strDigit))
			sum += current
		}
		ge.Done()
	}(digitChan)
	ge.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(sum)
}


// func main(){

// 	posibleLastDigits := [4]int{2, 4, 8, 6}
// 	lastDigitIndex := 0
// 	exp := 1

// 	for exp <= 1000 {
// 		fmt.Printf("El ultimo digito de 2^%v es %v.\n", exp, posibleLastDigits[lastDigitIndex])
// 		exp += 1
// 		lastDigitIndex += 1
// 		if lastDigitIndex > 3 {
// 			lastDigitIndex = 0
// 		}
// 	}

// }
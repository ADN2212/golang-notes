package main

import (
	"fmt"
	"sync"
)


var wg = sync.WaitGroup{}


func main(){
	//fmt.Println("Working")
	//Channels:
	//Estos son la forma en la que se usa el paralelismo en Go
	ch := make(chan int)//Un canal de enteros.
	wg.Add(2)//Agregar dos gorutines al grupo de espera.
	go func() {
		i := <- ch//Asi se resive data del chanel
		fmt.Println(i)
		ch <- 27
		wg.Done()//Da la gorutine por finalizada.
	}()//Porque una IIFE?
	go func(){
		//i := 52
		ch <- 52//Asi se envia info al canal
		//i = 100//Esta cambio no se ve en la primera gorutine.
		fmt.Println(<- ch)//El valor que esta enviando el canal en este momento.
		wg.Done()

	// }()
	// wg.Wait()//Espera a que termine la ejecucion de todas las Gorutines del grupo de espera.
	
	//Es posible que se tengan varias Gorutines enviando info a un solo canal como en el ejemplo siguiente:

	// ch := make(chan int)

	// for j := 0; j < 5; j++ {
	// 	wg.Add(2)
	// 	//Si esta gorutine estuviera afuera del ciclo el prograna se rompe
	// 	//porque al final de su ejecucion se cierra
	// 	//lo que haria que la segunda este enviando data que no va a ser resivida.
	// 	go func() {
	// 		i := <- ch
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// 	go func(){
	// 		ch <- 42
	// 		wg.Done()
	// 	}()	
	// }

	// wg.Wait()
	//fmt.Println(ch)

	//Las Gorutines pueden enviar y resivir datos de los canales
	//pero es posible hacer que unas solo resivan y otras solo envien como siguie:
	// ch := make(chan int)
	// wg.Add(2)
	// //Resive info de un canal de enteros:
	// go func (ch <- chan int){
	// 	i := <- ch
	// 	fmt.Printf("Resive el valor %v\n", i)
	// 	//ch <- 27//Must cause problems, no se puede enviar info en un canal que solo resive.
	// 	wg.Done()
	// }(ch)
	// //Envia info a un canal de enteros:
	// go func (ch chan <- int){
	// 	toSend := 42
	// 	ch <- toSend
	// 	fmt.Printf("Envia el valor %v\n", toSend)
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	//Buffered Channels:
	//Es posible que en alguna situacion se quira enviar mas de un dato a
	//un canal, para esto usaremos los Buffered chanels
	ch := make(chan int, 5)//El segundo parametro indica la cantidad de valores que pueden pasar a traves del canal. 
	wg.Add(2)
	//Resive info de un canal de enteros:
	go func (ch <- chan int){
		//Aqui podriamos ciclar sobre todos los valores que se envian a travez del canal.
		// i := <- ch//Primer valor.
		// fmt.Printf("Resive el valor %v\n", i)
		// i = <- ch//Segundo.
		// fmt.Printf("Resive el valor %v\n", i)

		//Se puede iterar a travez de los datos que estan en un canal.
		//Dado que la cantidad de datos que pueden pasar por un canal es indeterminada
		//Para poder usar el 'for-range' el canal debe cerrarce antes.
		for val := range ch {
			fmt.Printf("%v ha sido resivido...\n", val)
		}
		
		//Si no sabemos si el canal esta cerrado podemos interar sobre el hasta que no devuelva valores como sigue:

		// for {
		// 	//Aqui, ok sera false cuando no queden valores en el canal.
		// 	if val, ok := <- ch; ok{
		// 		fmt.Printf("%v ha sido resivido...\n", val)
		// 	} else {
		// 		break
		// 	}
		// }

		wg.Done()
		
	}(ch)
	// go func (ch <- chan int){
	// 	i := <- ch
	// 	fmt.Printf("Resive el valor %v\n", i)
	// 	wg.Done()
	// }(ch)
	//Envia info a un canal de enteros:
	go func (ch chan <- int){
		//toSend := 42
		ch <- 42
		ch <- 27//Trow an error, si el canal no soporta mas de un dato
		ch <- 45
		ch <- 3
		ch <- 5
		close(ch)//Cierra el canal y evita errores en el for-range.
 		//Ojo: No se pueden enviar datos a travez de un canal cerrado.
 		wg.Done()
	}(ch)

	wg.Wait()

 }


//Pendiente hacer un ejemplo del uso de los select statements..
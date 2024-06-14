package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"//Sirve para sincronizar Goroutines
)

func main() {
	//showGorutine(4)

	inicio := time.Now()

	grupoEspera := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		//Forma Secuencial:
		//Aqui estamos ejecutando cada una de nuestras Gorutines ficticias
		//La gorutine i + 1 debe espera a la finalizacion de la i para empezar (ejecucion secuencial).
		//showGorutine(i)

		//Forma concurrente:
		grupoEspera.Add(1)//Agregar una gorutine al grupo de espera.
		go showGorutine(i, &grupoEspera)
		//Vease como al final de la ejecucion las gorutines no estan ordenadas, esto es por la concunrrencia.
		//es decir, cada calculo se ha hecho en diferentes gorutines y ha finalizado en funcion del tiempo de espera.

	}

	//Ojo, si la funcion main termina se ejucion todas las demas gorutines tambien terminan.
	grupoEspera.Wait()//Espera a que se termine la ejecucion de todas las gorutines que tiene dentro.
	tiempoDeEjecucion := time.Since(inicio).Milliseconds()
	//fmt.Printf("La ejecución ha finalizado en %v nanosegundos.", tiempoDeEjecucion)
	fmt.Printf("La ejecución ha terminado en %v millisengundos.\n", tiempoDeEjecucion)
	//Observese la diferencia en el tiempo de ejecucion al quitar la palabra 'go'
}

//Aqui estamos pasando el puntero del grupo de espera para poder modificarlo desde dentro de la func
func showGorutine(id int, gePointer *sync.WaitGroup){
	lapse := time.Duration(rand.Intn(500))//Genera un numero aleatorio entre 0 y el argumento.
	//fmt.Println(lapse)
	time.Sleep(lapse * time.Millisecond)//Una parada de 100 mili segundos.
	(*gePointer).Done()//Ejecutamos este metodo para decir que esta gorutine esta terminada.
	fmt.Printf("Gorutine No. %v terminada.\n", id)
	//fmt.Printf("Gorutine No. %v ejecutada en %v nanosegundos.\n", id, lapse.Nanoseconds())
}


//Ojo: este es uno de los temas mas complejos de Go, por lo que hay que seguir investigando.





package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func jugador(nombre string, juego chan int) {
	defer wg.Done()

	for {
		//¿Hay pelota en el juego? (Se cerró el canal)
		pelota, existe := <-juego
		if !existe {
			//Se gana al estar cerrado el canal
			fmt.Printf("¡%s ganó!\n", nombre)
			return
		}

		//Escogemos un número al azar, será la suerte del jugador
		suerte := rand.Intn(100)
		if suerte%13 == 0 {
			fmt.Printf("%s perdió en la ronda: %d\n", nombre, pelota)
			//Para decir que perdió
			close(juego)
			return
		}

		//Incrementamos el conteo
		fmt.Printf("%s pegó en la ronda %d\n", nombre, pelota)
		pelota++
		//Mandamos la pelota al otro jugador
		juego <- pelota
	}
}

func main() {

	juego := make(chan int)

	wg.Add(2)

	go jugador("Nadal", juego)
	go jugador("Alberto", juego)

	//Lanzamos la pelota
	juego <- 1

	//Esperamos a que termine el juego
	wg.Wait()
}

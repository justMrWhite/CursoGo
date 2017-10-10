package main

import (
	"fmt"
	"sync"
)

var espera sync.WaitGroup
var candado sync.Mutex
var conteodeciclos int

func main() {

	for i := 1; i <= 10; i++ {
		espera.Add(1)
		go ciclointerno()
	}
	espera.Wait()
	fmt.Println("El conteo final es:", conteodeciclos)
}

func ciclointerno() {
	for i := 1; i <= 10; i++ {
		candado.Lock()
		x := conteodeciclos
		x++
		conteodeciclos = x
		candado.Unlock()
		//conteodeciclos++
	}
	espera.Done()
}

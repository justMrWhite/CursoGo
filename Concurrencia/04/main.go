package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
		//con esto ya no puedes meter más valores
		//sólo recibes el que le queda
	}()

	//Deja de hacer range hasta que el canal se cierra
	for v := range ch {
		fmt.Println(v)
	}

}

package main

import "fmt"

func main() {
	suma(3, 5, 10, 20, 15)
}

func suma(numeros ...int) {
	var contador int
	for _, v := range numeros {
		contador += v
	}
	fmt.Println(contador)
}

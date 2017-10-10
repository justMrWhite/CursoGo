package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	fmt.Println(len(numeros))
	fmt.Println(cap(numeros))

	numer := numeros[1:3]
	fmt.Println(numer)
	fmt.Println(len(numer))
	fmt.Println(cap(numer))

	fmt.Println("Reasignando")
	numer[0] = 100
	fmt.Println(numeros)
	fmt.Println(numer)

	/**
	rebanada[a:b] su slice tiene una capacidad de c
	longitud:    b - a
	capacidad:   c - a
	**/
}

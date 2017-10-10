package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	//numeros[0] = 100
	//fmt.Println(numeros)
	//numeros[5] = 1000 (error: fuera de rango)
	//fmt.Println(numeros)

	//append(a esto le metemos, esto)
	numeros = append(numeros, 1000)
	fmt.Println(numeros)
	fmt.Println(cap(numeros))
	fmt.Println(len(numeros))
}

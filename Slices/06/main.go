package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	for _, valores := range numeros {
		fmt.Println(valores)
	}
}

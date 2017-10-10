package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30

	//Ambas condiciones deben cumplirse
	fmt.Println("El operador 'y'")
	fmt.Println(a < b && a < c)
	fmt.Println(a < b && a > c)

	//Al menos una de las condiciones debe cumplirse
	fmt.Println("El operador 'o'")
	fmt.Println(a < b || a < c)
	fmt.Println(a < b || a > c)

	//El operador "no", cambia la condición
	fmt.Println("El operador 'no'")
	fmt.Println(!(a < c))
	fmt.Println(!true)

}

package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := 20
	fmt.Println("¿Son iguales?")
	fmt.Println(a == b)
	fmt.Println(a == c)

	fmt.Println("¿Son distintos?")
	fmt.Println(a != b)
	fmt.Println(a != c)

	fmt.Println("¿Es mayor o menor?")
	fmt.Println(a > c)
	fmt.Println(a < c)

	fmt.Println("¿Es mayor (menor) o igual?")
	fmt.Println(a >= b)
	fmt.Println(a <= c)
}

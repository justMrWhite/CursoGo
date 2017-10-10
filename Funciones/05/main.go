package main

import "fmt"

//var letra = "A"
//Esto es para ver el scope (alcance)

func main() {
	xy := []int{3, 5, 10, 20, 15}
	suma(xy...)
}

//func Suma(numeros ...int) { <- exported
func suma(numeros ...int) {
	var contador int
	for _, v := range numeros {
		contador += v
	}
	fmt.Println(contador)
	//fmt.Println(letra)
}

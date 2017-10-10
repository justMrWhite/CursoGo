package main

import "fmt"

func main() {
	x, y := retorno()
	fmt.Println(y)
	fmt.Println(x)
}

func retorno() (int, int) {
	return 10, 20
}

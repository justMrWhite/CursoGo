package main

import "fmt"

func main() {

	a := 7
	b := 14.0

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)

	//fmt.Println(a + b)
	//La anterior da un error por
	//tratarse de tipos diferentes

	fmt.Println(float64(a) + b)
}

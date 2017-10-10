package main

import "fmt"

func main() {
	a := 20
	fmt.Println(a)
	//Memory address
	fmt.Println(&a)

	a = 10
	fmt.Println(a)
	fmt.Println(&a)
}

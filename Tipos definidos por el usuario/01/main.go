package main

import "fmt"

func main() {
	type personalizado int

	var a int
	var b int
	var c personalizado

	a = 10
	b = 20
	c = 30

	fmt.Println(a + b)
	fmt.Println(c)
	fmt.Println(a + int(c))
}

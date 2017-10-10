package main

import "fmt"

func main() {
	frutas := []string{"mango", "sandía"}
	fmt.Println(frutas)

	fmt.Println(frutas[1])

	frutas[0] = "tuna"
	fmt.Println(frutas)
}

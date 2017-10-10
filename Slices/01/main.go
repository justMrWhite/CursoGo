package main

import "fmt"

func main() {
	letras := []string{"x", "y", "z"}
	fmt.Println(letras)
	//fmt.Println(len(letras))
	//fmt.Println(cap(letras))

	frutas := make([]string, 5)
	//frutas := make([]string, 5, 10)
	fmt.Println(frutas)
}

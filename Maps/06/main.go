package main

import "fmt"

func main() {

	diccionario := map[string]string{
		"A": "Primera letra del alfabeto",
		"B": "Segunda letra del alfabeto",
		"C": "Tercera letra del alfabeto",
	}

	fmt.Println(diccionario)
	// delete(map, key)
	delete(diccionario, "C")

	fmt.Println(diccionario)

}

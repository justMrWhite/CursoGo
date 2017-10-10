package main

import "fmt"

func main() {

	diccionario := map[string]string{
		"A": "Primera letra del alfabeto",
		"B": "Segunda letra del alfabeto",
		"C": "Tercera letra del alfabeto",
	}

	for clave, valor := range diccionario {
		fmt.Println(clave, valor)
	}

}

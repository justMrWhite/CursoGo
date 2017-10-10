package main

import "fmt"

func main() {

	diccionario := map[string]string{
		"A": "Primera letra del alfabeto",
		"B": "Segunda letra del alfabeto",
		"C": "Tercera letra del alfabeto",
	}
	//diccionario["D"] = "Cuarta letra del alfabeto"

	valor, existencia := diccionario["D"]
	if existencia {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe ese valor")
	}

	//fmt.Println(diccionario)
}

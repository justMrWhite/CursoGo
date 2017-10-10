package main

import "fmt"

func main() {

	var diccionario map[string]string
	//este mapa es nil

	diccionario["A"] = "Primera letra del alfabeto"
	//por ser nil, no puede aceptar entradas
	fmt.Println(diccionario)
}

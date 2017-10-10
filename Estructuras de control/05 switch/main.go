package main

import "fmt"

func main() {
	calificacion := "b"

	switch calificacion {
	case "m":
		fmt.Println("Nos diste un: malo")
	case "b":
		fmt.Println("Nos diste un: bueno")
	case "e":
		fmt.Println("Nos diste un: excelente")
	default:
		fmt.Println("Algo anda mal")
	}
	/**
		if calificacion == "m" {
			fmt.Println("Nos diste un: malo")
		} else if calificacion == "b" {
			fmt.Println("Nos diste un: bueno")
		} else if calificacion == "e" {
			fmt.Println("Nos diste un: excelente")
		} else {
			fmt.Println("Algo anda mal")
		}
	**/
}

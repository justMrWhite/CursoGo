package main

import "fmt"

type estudiante struct {
	nombre string
	correo string
}

func main() {
	a := estudiante{"Alberto", "a@correo.com"}
	fmt.Println(a)

	//Acceder:
	fmt.Println(a.correo)
	//Asignar:
	a.correo = "alberto@correo.com"
	fmt.Println(a.correo)
}

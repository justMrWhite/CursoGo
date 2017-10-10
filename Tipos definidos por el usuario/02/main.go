package main

import "fmt"

type alumno struct {
	nombre string
	edad   int
	correo string
}

func main() {
	a := alumno{"Alberto", 34, "a@correo.com"}
	m := alumno{
		nombre: "Mariana",
		edad:   32,
		correo: "m@correo.com",
	}

	fmt.Println(a)
	fmt.Println(m)
}

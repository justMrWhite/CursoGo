package main

import "fmt"

type estudiante struct {
	nombre string
	edad   int
	correo string
}

//Método (receiver)
func (e estudiante) damedatos() {
	fmt.Printf("Se llama %v, tiene %v años y su correo es %v \n",
		e.nombre, e.edad, e.correo)
}

func main() {
	ab := estudiante{"Alberto",
		34,
		"ab@correodelaescuela.com"}

	md := estudiante{"Mayra",
		32,
		"may@correodelaescuela",
	}

	//Llamar método sobre estudiante
	ab.damedatos()
	md.damedatos()
}

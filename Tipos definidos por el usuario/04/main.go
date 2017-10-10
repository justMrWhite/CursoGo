package main

import "fmt"

type estudiante struct {
	nombre string
	edad   int
	correo string
}

type jefedegrupo struct {
	//estudiante estudiante
	estudiante
	privilegios bool
}

//Type embedding

func main() {
	a := estudiante{"Alberto", 34, "correo@escuela.com"}
	fmt.Println(a)

	m := jefedegrupo{
		estudiante: estudiante{
			nombre: "Mariana",
			edad:   32,
			correo: "m@correo.com",
		},
		privilegios: true,
	}
	fmt.Println(m)
	fmt.Println(m.nombre)
	fmt.Println(m.estudiante.nombre)
}

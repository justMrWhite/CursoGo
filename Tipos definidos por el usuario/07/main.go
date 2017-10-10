package main

import "fmt"

type estudiante struct {
	nombre string
	correo string
}

func main() {
	ab := estudiante{"Alberto", "ab@escuela.com"}
	fmt.Println(ab.correo)

	ab.correo = "alberto@escuela.com"
	fmt.Println(ab.correo)

	ab.cambiacorreo("a@escuela.com")
	fmt.Println(ab.correo)
}

func (e *estudiante) cambiacorreo(nuevocorreo string) {
	e.correo = nuevocorreo
}

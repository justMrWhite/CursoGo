package main

import (
	"fmt"
	"math"
)

//Interface
type operador interface {
	perimetro() float64
	area() float64
}

//Función de interface
func operar(o operador) {
	fmt.Println(o.perimetro())
	fmt.Println(o.area())
}

type cuadrado struct {
	lado float64
}

type circulo struct {
	radio float64
}

func (cu cuadrado) perimetro() float64 {
	return cu.lado * 4
}

func (ci circulo) perimetro() float64 {
	return 2 * math.Pi * ci.radio
}

func (cu cuadrado) area() float64 {
	return cu.lado * cu.lado
}

func (ci circulo) area() float64 {
	return math.Pi * ci.radio * ci.radio
}
func main() {
	redondin := circulo{7}
	cuadradin := cuadrado{10}

	//Llamar como interface
	operar(redondin)
	operar(cuadradin)

	/**
		fmt.Println(redondin.perimetro())
		fmt.Println(redondin.area())
		fmt.Println(cuadradin.perimetro())
		fmt.Println(cuadradin.area())
	**/
}

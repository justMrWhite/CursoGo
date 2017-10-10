package main

import "fmt"

func main() {

	var nombre string
	nombre = "Alberto"

	var edad int
	edad = 34

	var estatura float64
	estatura = 1.83

	var inteligente bool
	inteligente = true

	fmt.Println(nombre, edad, estatura, inteligente)
}

/**
TIPOS DE DATOS BÁSICOS - BASIC TYPES
bool -> true o false (verdadero o falso)

string -> "texto"

int  int8  int16  int32  int64 -> números enteros
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64 -> números no enteros

complex64 complex128
**/

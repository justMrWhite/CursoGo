package main

import "fmt"

func main() {
	calificaciones := []float64{7.6, 9, 7.9, 8.4, 9.1}
	promedio(calificaciones)
}

func promedio(cal []float64) {
	var suma float64
	for _, valor := range cal {
		suma += valor
	}
	fmt.Println(suma / float64(len(cal)))
}

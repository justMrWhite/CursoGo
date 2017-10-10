package main

import "fmt"

func main() {
	calificaciones := []float64{7.6, 9, 7.9, 8.4, 9.1}
	fmt.Println(promedio(calificaciones))
}

func promedio(cal []float64) float64 {
	//suma := 0.0
	var suma float64
	for _, valor := range cal {
		//suma = suma + valor
		suma += valor
	}
	return suma / float64(len(cal))
}

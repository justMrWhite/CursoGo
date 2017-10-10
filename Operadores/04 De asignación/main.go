package main

import "fmt"

func main() {
	x := 10

	x = 20
	fmt.Println("Operador de asignación:", x)

	//x = x + 10
	x += 10
	fmt.Println("A la misma variable sumamos 10:", x)

	//x = x - 10
	x -= 10
	fmt.Println("A la misma variable restamos 10:", x)

	//La misma lógica aplica con * % /
	//x = x * 2
	x *= 2
	fmt.Println("La misma variable la multiplicamos por 2:", x)
}

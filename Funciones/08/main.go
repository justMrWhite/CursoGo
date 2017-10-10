package main

import "fmt"

func main() {
	uno := 10
	dos := 10

	suma(uno, dos)
	defer resta(uno, dos)
	multi(uno, dos)
}

func suma(x, y int) {
	fmt.Println(x + y)
}

func resta(x, y int) {
	fmt.Println(x - y)
}

func multi(x, y int) {
	fmt.Println(x * y)
}

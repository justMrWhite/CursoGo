package main

import "fmt"

func main() {
	compartido := 0

	agregauno := func() int {
		compartido++
		return compartido
	}

	fmt.Println(agregauno())
	fmt.Println(agregauno())
	fmt.Println(agregauno())
}

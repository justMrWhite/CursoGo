package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	edad := 17

	if edad < 14 {
		fmt.Println("Estás muy pequeño")
	} else if edad == 15 {
		fmt.Println("Ya casi")
	} else {
		fmt.Println("Ya la hiciste")
	}
}

package main

import "fmt"

func main() {
	var slice []int
	fmt.Println(slice)
	if slice == nil {
		fmt.Println("nil")
	}

	slice2 := make([]int, 0)
	fmt.Println(slice2)
	if slice2 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("Tu slice no es nil")
	}
}

package main

import "fmt"

func main() {
	fmt.Println(recursiva(3))
}

func recursiva(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursiva(n-1)
}

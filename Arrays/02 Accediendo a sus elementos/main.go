package main

import "fmt"

func main() {

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	var array3 [5]int
	array3 = array2
	fmt.Println(array3)
	//array2[4] = 7
	//fmt.Println(array2)
	//fmt.Println(array2[2])

}

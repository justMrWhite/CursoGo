package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go esto()
	go lootro()
	wg.Wait()
}

func esto() {
	for i := 1; i <= 50; i++ {
		fmt.Println(i, "esto")
		time.Sleep(1 * time.Millisecond)
	}
	wg.Done()
}

func lootro() {
	for i := 1; i <= 50; i++ {
		fmt.Println(i, "lo otro")
		time.Sleep(2 * time.Millisecond)
	}
	wg.Done()
}

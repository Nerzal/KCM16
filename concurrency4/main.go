package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var foobar int
	foobar = 3
	fmt.Println(foobar)

	caribar := 3
	fmt.Println(caribar)
	wg.Add(3)
	go foo(1, &wg)
	go foo(2, &wg)
	go foo(3, &wg)

	wg.Wait()
	fmt.Println("Program Finished")
}

func foo(number int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(number)
	wg.Done()
}

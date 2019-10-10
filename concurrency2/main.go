package main

import (
	"fmt"
	"time"
)

func main() {
	go foo(1)
	go foo(2)
	go foo(3)
	fmt.Println("Program Finished")
}

func foo(number int) {
	time.Sleep(1 * time.Second)
	fmt.Println(number, "finished")
}

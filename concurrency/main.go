package main

import (
	"fmt"
	"time"
)

func main() {
	foo(1)
	foo(2)
	foo(3)
	fmt.Println("Program Finished")
}

func foo(number int) {
	time.Sleep(1 * time.Second)
	fmt.Println(number, "finished")
}

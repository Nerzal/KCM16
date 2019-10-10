package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 3)

	go foo(1, channel)
	go foo(2, channel)
	go foo(3, channel)

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Program Finished")
}

func foo(number int, channel chan int) {
	time.Sleep(1 * time.Second)

	channel <- number
}

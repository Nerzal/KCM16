package main

import "fmt"

type MyStruct struct {
	Foo string
}

type privateStruct struct {
}

func (mystruct MyStruct) FooBar() error {
	return nil
}

func main() {
	//FooBar()

	myStruct := &MyStruct{}
	err := myStruct.FooBar()
	if err != nil {
		fmt.Println(err)
	}

	myStruct2 := MyStruct{}
	err = myStruct2.FooBar()
	if err != nil {
		fmt.Println(err)
	}
}

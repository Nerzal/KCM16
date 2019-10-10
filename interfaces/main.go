package main

import "fmt"

// MyInterface is used as implementation
type MyInterface interface {
	Foo() error
}

type implementation struct {
}

func (imp implementation) Foo() error {
	return nil
}

// NewMyInterface instantiates a new MyInterface
func NewMyInterface() MyInterface {
	return &implementation{}
}

func main() {
	implementation := NewMyInterface()
	err := implementation.Foo()
	if err != nil {
		fmt.Println(err)
	}

	FooBar(implementation)
}

func FooBar(baz MyInterface) {
	err := baz.Foo()
	if err != nil {
		fmt.Println(err)
	}
}

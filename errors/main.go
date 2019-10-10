package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	foo()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	if err := bar(); err != nil {
		fmt.Println(err)
	}
}

func foo() (int, error) {
	return 0, errors.New("foo failed")
}

func bar() error {
	err := foo()
	return errors.Wrap(err, "Bar failed")
}

func baz() error {
	return &MyError{}
}

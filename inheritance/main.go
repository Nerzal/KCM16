package main

import "fmt"

func main() {
	foo := Foo{Name: "Alice"}
	fmt.Println(foo)

	bar := Bar{
		Foo: Foo{
			Name: "Bob",
		},
	}
	fmt.Println(bar)

	baz := Baz{
		Foo: Foo{
			Name: "Bob",
		},
	}
	fmt.Println(baz)
}

type Foo struct {
	Name string
}

type Bar struct {
	Foo
}

type Baz struct {
	Foo Foo
}

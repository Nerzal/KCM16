package main

type MyError struct {
}

func (err *MyError) Error() string {
	return "MyError"
}

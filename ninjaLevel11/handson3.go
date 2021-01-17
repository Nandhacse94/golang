package main

import (
	"fmt"
	"errors"
)

type customError struct {
	info string
	code int
	err  error
}

func (e customError) Error() string {
	return fmt.Sprintf("here is the error %v -[%v], base:%v",e.info,e.code, e.err)
}

func foo(c customError) {
	fmt.Println("In foo: ", c)
}

func main() {
	//custom error calls which implements interface.

	c := customError{
		info: "Needs coffee",
		code: 201,
		err:  errors.New("Boring"),
	}

	foo(c)
}

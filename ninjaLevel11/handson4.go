package main

import (
	"errors"
	"fmt"
)

type sqrtError struct {
	info string
	code int
	err  error
}

func (e sqrtError) Error() string {
	return fmt.Sprintf("here is the sqrtError %v -[%v], base:%v", e.info, e.code, e.err)
}

func foo(c sqrtError) {
	fmt.Println("In foo: ", c)
}

func main() {
	//custom error calls which implements interface.

	_, err := sqrt(-10.23)
	if err != nil {
		fmt.Println("Err:", err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{"Needs coffee", 201, errors.New("Boring")}
	}
	return f, nil
}


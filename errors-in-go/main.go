package main

import (
	"errors"
	"fmt"
)

var (
	ErrIslessThanZero = errors.New("Number less than zero")
	ErrIsNotEven      = errors.New("Number is not even")
)

func IsEven(num int) error {
	if num&2 != 0 {
		return ErrIsNotEven
	}
	return nil
}

func main() {
	fmt.Println("Errors in Golang")

	err := IsEven(23)
	if err != nil {
		fmt.Println("Number is not even same shit : err code: ", err.Error())
	}
}

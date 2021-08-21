package main

import (
	"errors"
	"fmt"
)

func IPanicEasily() {
	defer func() {
		fmt.Println("I am deffered")
		fmt.Println("1")

	}()
	panic("I panic easily")
}

func myAwesomeFunction() (err error) {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered from a small panic")
			err = errors.New("I panic easily panicked")
		}

		// eğer bu recover fonksiyonu çalışmazsa maindeki diğer işlemlere devam edemeyiz"  main function successfully "  basamayız

		/* 	fmt.Println("my awesome function deffered")
		fmt.Println("2") */
	}()

	IPanicEasily()
	return nil
}

func main() {
	/*
		fmt.Println("Paniccccc!! yok dayı yazlımcı burda")
		fmt.Println("Panic in the go ap")

		panic("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAH!!!!")
	*/

	defer func() {
		fmt.Println("Main function deffered")
		fmt.Println("3")
	}()

	if err := myAwesomeFunction(); err != nil {
		fmt.Println("my awesome function returned error.")
		fmt.Println(err.Error())
	}
	fmt.Println("main function successfully ")
}

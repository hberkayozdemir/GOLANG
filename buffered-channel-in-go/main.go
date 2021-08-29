package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValues(values chan int) {

	for i := 0; i <= 10; i++ {

		value := rand.Intn(10)
		fmt.Printf("CalculateValue %d \n", value)
		values <- value
	}
}

func main() {
	values := make(chan int, 4)
	go CalculateValues(values)

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		value := <-values

		fmt.Println(value)

	}
}

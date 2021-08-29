package main

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Printf("Calculated Value %d", value)
	values <- value
}

func main() {

	values := make(chan int)
	go CalculateValue(values)
	value := <-values
	fmt.Println(value)

}

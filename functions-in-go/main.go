package main

import "fmt"

func HelloWorld(name string, age, height int, pi, periscope float64) {

	fmt.Println("Hello  ", name)
	fmt.Println("Age ", age)
	fmt.Println("Height ", height)
	fmt.Println(pi * periscope)
}

func addTotal(a, b int) (int, int) {
	return a + b, a - b
}

// errorları döndürme burdan geliyor heralde
func main() {
	HelloWorld("Hilmi", 23, 190, 3.14, 12332)

	total, minus := addTotal(2, 3)

	fmt.Println(total)
	fmt.Println(minus)
}

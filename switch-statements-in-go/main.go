package main

import (
	"runtime"
)

func main() {
	var customerHeight int = 140
	customerAge := 18

	switch {

	case customerHeight >= 150 || customerAge < 180:
		println("Can acces ride")
	case customerHeight >= 120:
		println("Can acces children's rides")
	default:
		println("can not acces rides")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		println("os x")
	case "linux":
		println("linux machine")
	default:
		println("something else...")
	}

}

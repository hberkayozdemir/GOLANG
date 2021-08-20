package main

func main() {

	println("If statements in Go")

	var customerHeight int = 160 // if
	/*  var customerHeight int = 140 // else if */
	/* 		var customerHeight int = 110 // else  */

	customerAge := 18

	if customerHeight >= 150 && customerAge >= 18 {
		println("Can access ride")
	} else if customerHeight >= 120 {

		println("Customer can access chidrens's rides")
	} else {
		println("Customer is too small")
	}

}

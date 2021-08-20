package main

import "fmt"

func main() {

	ages := []int{42, 28, 30, 31, 27, 18}
	fmt.Println("##################################################################################")
	for index, value := range ages {

		fmt.Println(index, " index of ages array value is = ", value)

	}
	fmt.Println("##################################################################################")
	fmt.Println("Range Yukarısı ")

	for i := 0; i < len(ages); i++ {

		fmt.Println(ages[i])
	}

	for i := 0; i < len(ages); {
		fmt.Println(ages[i])
		i++
	}

	for {
		fmt.Println("I started and finish")
		break
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

}

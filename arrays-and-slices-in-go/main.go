package main

import "fmt"

// format package
func main() {

	println("Arrays and Slices in Go")

	planets := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}

	// println() fonskiyponu çalışmaz array basmak için dolayısı ile fmt yi import edip şu şekilde basacağız
	fmt.Println(planets)

	var planetsArray [8]string
	planetsArray[0] = "mercury"

	fmt.Println(planetsArray)
	// empty element leri boş basıyor

	planetSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(planetSlice)

	var planetSliceVerbose []string
	planetSliceVerbose = append(planetSliceVerbose, "Mercury")
	// append edildiğinde otomatik resizing logic çalışıyor
	fmt.Println(planetSliceVerbose)
	// slice ta size belirtmek zorunda değiliz eklediğimiz kadar go compiler otomatik olarak size
	// bellirliyor
}

package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}
type Project struct {
	Name        string
	Priortiy    string
	Technolgies []string
}

func main() {
	fmt.Println("             Struct      ")
	fmt.Println("##############################")

	engineer := Engineer{
		Name: "hilmi berkay",
		Age:  28,
		Project: Project{
			Name:        "Begginner's guide to Go",
			Priortiy:    "High",
			Technolgies: []string{"Go"},
		},
	}
	// amma şu şekilde tanımlayadabiliriz
	/* 	engineer := Engineer{"hilmi berkay", 28} */
	fmt.Println(engineer)
	// bu şekilde getirirsem sadece isimin değerini basıyor ama structtaki değeride görmek istiyorsak bunu değşştirmek icin sunu kullanmalıyız
	fmt.Println("================================================================")
	fmt.Printf("%+v\n", engineer)
	fmt.Println(engineer.Project.Name)

}

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

func (e Engineer) Print() {

	fmt.Println("Engineer Information")
	fmt.Printf("Name : %s\n", e.Name)
	fmt.Printf("Age %d\n", e.Age)
	fmt.Printf("Current Project : : %s\n", e.Project.Name)

}

//func (e Engineer) updateAge() { bu şekilde kullanırsan değerde değişiklik yapamazsın
func (e *Engineer) updateAge() {
	e.Age += 1

}

// return value string bunda şart değil ama koyarsakta fark etmez en iyisi koymak hepsine :D
func (e Engineer) GetProjectPriority() string {

	return e.Project.Priortiy
}

func main() {

	engineer := Engineer{
		Name: "hilmi berkay",
		Age:  28,
		Project: Project{
			Name:        "Begginner's guide to Go",
			Priortiy:    "High",
			Technolgies: []string{"Go"},
		},
	}
	fmt.Println(engineer)
	// şimdi biz update Age te enginner ın basına receiver pointer koymadıgımızdad eğişklik yapamıyoreuz değşştirmek
	// yaş 28 olarak tanumlıydı ve artması gerekirken artmadı fakat pointer tanımlarsak arttırabiliyouz
	// value larda değişiklik yapmak istiyorsak "*" kullanmamız şart
	engineer.Print()
	engineer.updateAge()
	engineer.Print()
	fmt.Println(engineer.GetProjectPriority())
}

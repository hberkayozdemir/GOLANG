// BASIC EXAMPLE
/* package main

import "fmt"

func main() {

	fmt.Println("Go pointers tutorials")

	var name string
	name = "elliot"
	fmt.Println(name)

	ptr := &name
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = "donna"
	fmt.Println(name)
}
*/

// COMPLEX EXAMPLE

package main

import "fmt"

func (e *Engineer) UpdateAge() {
	e.Age += 1
}

// böyle yaşı değiştirdik
func (e Engineer) UpdateName(name string) {

	e.Name = name
}
func (e *Engineer) UpdateNameWithPointer(name string) {

	e.Name = name
}

type Engineer struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Go pointers tutorial")

	elliot := &Engineer{Name: "elliot", Age: 27}
	fmt.Println(elliot)
	elliot.UpdateAge()
	fmt.Println(elliot)
	elliot.UpdateName("Hüsnüü")
	fmt.Println(elliot)
	fmt.Println(" Burda pointer ekleniyor || here I added pointer to this function")
	elliot.UpdateNameWithPointer("Hüsnüü")
	fmt.Println(elliot)
}

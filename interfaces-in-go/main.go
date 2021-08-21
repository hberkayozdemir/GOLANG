package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}
type Worker struct {
	Name string
}
type Manager struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Engineer name is " + e.Name
}
func (w *Worker) GetName() string {
	return "Worker name is " + w.Name
}

func (m *Manager) GetName() string {
	return "Manager name is " + m.Name
}
func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}
func main() {
	// eğer pointer receiverı koyuyorsak engineer ataması yaparken & kullanmak zorundayız fakat biz burda pointer kullanmak zorunda değildik dolayısı ile
	// pointerı kaldırırsan & bunuda kaldırabiliriz pointerı biz değişiklik yapacağımız değerlere yazıyoruz
	engineer := &Engineer{Name: "Elliot"}
	PrintDetails(engineer)
	worker := &Worker{Name: "Kazım"}
	PrintDetails(worker)
	manager := &Manager{Name: "Doruk"}
	PrintDetails(manager)
}

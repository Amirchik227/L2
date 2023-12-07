package main

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
	Паттерн "Фабричный метод" (Factory Method) — это порождающий паттерн проектирования, который предоставляет интерфейс для создания экземпляра класса, но оставляет выбор подкласса, экземпляр которого будет создан, самому объекту-потребителю.

	Применимость:
	1) Когда заранее неизвестно, объекты каких классов нужно создавать.
	2) Когда система должна быть независимой от процесса создания новых объектов и расширяемой.

	Плюсы:
	1) Отделение создания объектов от их использования: Позволяет клиентскому коду использовать абстрактный интерфейс для создания объектов, избегая зависимости от конкретных классов.
	2) Легкость добавления новых продуктов: Новые продукты могут быть добавлены без изменения существующего клиентского кода, благодаря использованию абстрактных фабрик и интерфейсов.
	3) Поддержка принципа инверсии зависимостей: Клиент зависит от абстракции, а не от конкретной реализации, что способствует принципу инверсии зависимостей.

	Минусы:
	1) Усложнение структуры кода: Внедрение фабричных методов может привести к увеличению числа классов и усложнению структуры кода.
	2) Множество классов для каждого продукта: Если существует много различных продуктов, для каждого из которых требуется своя фабрика, это может привести к созданию большого числа подклассов фабрик и продуктов.
	3)Даже для одного объекта необходимо создать соответствующую фабрику, что увеличивает код.
*/

import (
	"fmt"
)

// Интерфейс телефона
type ISmartPhone interface {
	Info()
	Ring()
}

// Первый телефон
type Nokia struct {
	model   string
	display float32
}

// Функция вывода информации о телефоне
func (n Nokia) Info() {
	fmt.Printf("Model: NOKIA %s, Display:%.1f\n", n.model, n.display)
}

// Звонок на первый телефон
func (n Nokia) Ring() {
	fmt.Println("BZZ-BZZ-BZZ")
}

// Создание первого телефона
func NewNokia() ISmartPhone {
	return &Nokia{
		model:   "5228",
		display: 3.2,
	}
}

// Второй телефон
type Sony struct {
	model   string
	display float32
}

// Функция ввода информации о втором телефоне
func (s Sony) Info() {
	fmt.Printf("Model: SONY %s, Display:%.1f\n", s.model, s.display)
}

// Звонок на второй телефон
func (s Sony) Ring() {
	fmt.Println("Pam-Pam")
}

// Создание второго телефона
func NewSony() ISmartPhone {
	return &Sony{
		model:   "W810i",
		display: 1.9,
	}
}

// Фабрика по созданию телефонов
func CreatePhone(phoneType string) (ISmartPhone, error) {
	if phoneType == "Nokia" {
		return NewNokia(), nil
	}
	if phoneType == "Sony" {
		return NewSony(), nil
	}
	return nil, fmt.Errorf("Wrong phone type")
}

// func main() {
// 	phones := [...]string{"Nokia", "Sony"}

// 	for i := 0; i < 5; i++ {
// 		randomModel := phones[rand.Intn(len(phones))]
// 		phone, _ := CreatePhone(randomModel)
// 		phone.Info()
// 		phone.Ring()
// 	}
// }
/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

	Паттерн "Строитель" (Builder) является порождающим паттерном проектирования, который используется для пошагового конструирования сложного объекта. Он позволяет создавать различные представления конечного объекта, разделяя процесс конструирования и представление конечного продукта.

	Применимость:
	1) Когда процесс конструирования объекта состоит из множества шагов.
	2) Когда требуется создавать разные представления одного и того же объекта.
	3) Когда конструирование объекта должно быть независимым от его представления и использования.

	Плюсы:
	1) Шаг за шагом конструирование: Позволяет создавать объекты пошагово, контролируя каждый шаг процесса.
	2) Раздельное конструирование и представление: Позволяет использовать один и тот же строитель для создания различных представлений объекта.
	3) Упрощение клиентского кода: Клиентский код не зависит от конкретных классов продуктов или их конструирования, что делает его более гибким.
	4) Изолирует сложный код сборки объекта от его основной бизнес-логики.

	Минусы:
	1) Увеличение сложности кода: Внедрение строителя может привести к увеличению числа классов и сложности кода.
	2) Невозможность использования строителя внутри const-структур: Некоторые языки программирования, такие как Go, не позволяют использовать строитель внутри const-структур из-за ограничений на вызов не-const методов внутри них.
	3)Клиент может оказаться привязан к конкретным классам строителей, так как в интерфейсе строителя может не быть метода получения результата.
*/

package main

import "fmt"

// Структура телефонов
type Phone struct {
	chip    string
	display float64
	ram     uint
}

// Вывод характеристик  телефона
func (p Phone) Info() {
	fmt.Printf("Chip: %s, Display: %.1f, RAM: %d\n", p.chip, p.display, p.ram)
}

// Билдер для cамсунга
type SamsungBuilder struct {
	chip    string
	display float64
	ram     uint
}

func (s *SamsungBuilder) setChip() {
	s.chip = "Snapdragon"
}

func (s *SamsungBuilder) setDisplay() {
	s.display = 6.8
}

func (s *SamsungBuilder) setRAM() {
	s.ram = 12
}

// Создание телефона билдером
func (s *SamsungBuilder) getPhone() Phone {
	return Phone{
		chip:    s.chip,
		display: s.display,
		ram:     s.ram,
	}
}

//  Функция для создания билдера телефона samsung
func newSamsungBuilder() *SamsungBuilder {
	return &SamsungBuilder{}
}

// Билдер для Пикселя
type PixelBuilder struct {
	chip    string
	display float64
	ram     uint
}

func (s *PixelBuilder) setChip() {
	s.chip = "Tensor"
}

func (s *PixelBuilder) setDisplay() {
	s.display = 6.4
}

func (s *PixelBuilder) setRAM() {
	s.ram = 8
}

// Создание телефона билдером
func (s *PixelBuilder) getPhone() Phone {
	return Phone{
		chip:    s.chip,
		display: s.display,
		ram:     s.ram,
	}
}

//  Функция для создания билдера телефона pixel
func newPixelBuilder() *PixelBuilder {
	return &PixelBuilder{}
}

// Интерфейс билдера телефонов
type IBuilder interface {
	setChip()
	setDisplay()
	setRAM()
	getPhone() Phone
}

// Функция получения билдера для некоторого телефона
func getBuilder(builderType string) IBuilder {
	if builderType == "Samsung" {
		return newSamsungBuilder()
	}

	if builderType == "Pixel" {
		return newPixelBuilder()
	}
	return nil
}

// Интерфейс директора, управляющего билдерами
type Director struct {
	builder IBuilder
}

// Функция по созданию пового билдера телефонов
func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// Передача билдера директору
func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

// После того как директор получил билдера, он может создать телефон с его помощью
func (d *Director) buildPhone() Phone {
	d.builder.setChip()
	d.builder.setDisplay()
	d.builder.setRAM()
	return d.builder.getPhone()
}

// func main() {
// 	// Создаем билдер для самсунга
// 	samsungBuilder := getBuilder("Samsung")
// 	// Создаем директора для производства самсунгов
// 	director := newDirector(samsungBuilder)
// 	// Директор создает новый самсунг
// 	samsung := director.buildPhone()
// 	// Выводим информацию о новом телефоне
// 	samsung.Info()

// 	// Теперь создаем билдер для пикселя
// 	pixelBuilder := getBuilder("Pixel")
// 	// Переназначаем билдера у дикректора
// 	// Теперь он будет управлять производством пикселей
// 	director.setBuilder(pixelBuilder)
// 	// Создаем пиксель
// 	pixel := director.buildPhone()
// 	pixel.Info()
// }

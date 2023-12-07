package main

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

	Паттерн "Состояние" (State) — это поведенческий паттерн проектирования, который позволяет объекту изменять свое поведение в зависимости от своего внутреннего состояния. При этом объект выглядит так, будто изменяет свой класс.

	Применимость:
	1)Когда у объекта есть внутреннее состояние, которое влияет на его поведение, и это состояние может изменяться во время выполнения.
	2) Когда код объекта содержит множество условных операторов, зависящих от его текущего состояния.

	Плюсы:
	1) Четкое разделение ответственности: Каждое состояние инкапсулирует свое поведение, что облегчает добавление новых состояний и изменение поведения объекта.
	2) Устойчивость к изменениям: Паттерн позволяет добавлять новые состояния и изменять поведение объекта без изменения существующего кода.

	Минусы:
	1) Увеличение числа классов: Внедрение состояний может привести к увеличению числа классов в системе.
	2) Сложность для небольших объектов: Для небольших объектов, у которых состояний мало, использование этого паттерна может быть избыточным.
*/

// Интерфейс всевозможных сосотяний
type State interface {
	requestItem() error
	dispenseItem() error
}

// Торговый автомат
type VendingMachine struct {
	selectionAvailable State // Сосотояние выбора товара
	itemRequested      State // Состояние выдачи товара

	currentState State // Текщее состояние
}

// При выборе товара вызоваем соответсвующий обработчик у объекта
// соответствующего текущему состоянию
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

// Создание торгового автомата
func newVendingMachine() *VendingMachine {
	v := &VendingMachine{}

	selectItemState := &SelecetItemState{
		vendingMachine: v,
	}
	despenseItemState := &DespenseItemState{
		vendingMachine: v,
	}

	v.itemRequested = despenseItemState
	v.selectionAvailable = selectItemState

	v.currentState = selectItemState
	return v
}

// Состояние выбора товара
type SelecetItemState struct {
	vendingMachine *VendingMachine
}

// Покупатель запросил некоторый товар
func (s *SelecetItemState) requestItem() error {
	fmt.Println("Подготовка товара")
	// Смена состояния
	s.vendingMachine.currentState = s.vendingMachine.itemRequested
	return nil
}

// Выдача товара пользователю в данном состоянии невозможна,
// так как товар не выбран
func (s *SelecetItemState) dispenseItem() error {
	return errors.New("Сначала выберете товар")
}

// Состояние выдачи товара
type DespenseItemState struct {
	vendingMachine *VendingMachine
}

// Выбор товара во время состояния выдачи невозможен
func (s *DespenseItemState) requestItem() error {
	return errors.New("Подождите, происходит выдача товара")
}

// Запрос на выдачу товара
func (s *DespenseItemState) dispenseItem() error {
	fmt.Println("Выдача товара прошла успешно")
	// Смена состояния
	s.vendingMachine.currentState = s.vendingMachine.selectionAvailable
	return nil
}

func main() {
	//Создание торгового автомата
	vendingMachine := newVendingMachine()

	fmt.Println("Выберете товар")

	// Выбор товара
	err := vendingMachine.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Выдача товара
	err = vendingMachine.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()

	fmt.Println("Выберете товар")
	// Выбор товара
	err = vendingMachine.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Выдача товара
	err = vendingMachine.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}
}

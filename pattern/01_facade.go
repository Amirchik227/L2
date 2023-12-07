package main

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

	Паттерн "Фасад" (Facade) является структурным паттерном проектирования, который предоставляет унифицированный интерфейс к группе интерфейсов подсистемы. Это упрощает использование подсистемы, предоставляя более высокоуровневый интерфейс.
	Применимость:
	1)Когда нужно предоставить простой или унифицированный интерфейс к сложной системе.
	2)Когда необходимо объединить несколько интерфейсов в одном, чтобы сделать работу с системой более удобной.
	3)Когда нужно уменьшить зависимости между клиентом и сложной системой, разбивая их на отдельные компоненты.

	Плюсы:
	1) Упрощение интерфейса: Фасад предоставляет простой интерфейс для взаимодействия с подсистемой, скрывая ее сложность от клиента.
	2) Уменьшение зависимостей: Клиент зависит только от фасада, а не от подсистемы, что уменьшает связанность между клиентом и подсистемой.
	3) Повышение уровня абстракции: Фасад позволяет работать с подсистемой на более высоком уровне абстракции, скрывая детали реализации.

	Минусы:
	1) Ограниченность функциональности: Фасад может предоставлять ограниченный набор функций, что может быть недостаточным для некоторых клиентов.
	2) Дополнительная сложность: Внедрение фасада может добавить дополнительный уровень абстракции, что может усложнить систему.
*/

import (
	"errors"
	"math/rand"
)

// Подсистема, которая проверяет наличие свободных номеров
type AvailabilityChecker struct{}

func (a *AvailabilityChecker) checkAvailability() bool {
	return rand.Intn(100) > 30
}

// Подсистема, которая выбирает для вас конкертный номер в отеле
type RoomSelector struct{}

func (r *RoomSelector) selectRoom() int {
	return rand.Intn(1000)
}

// Подсистема, которая подтверждает резервирование номера
type ReservationSystem struct{}

func (rs *ReservationSystem) confirmReservation() bool {
	return true
}

// Фасад для бронироания номера в отеле
type HotelBookingFacade struct {
	availabilityChecker *AvailabilityChecker
	roomSelector        *RoomSelector
	reservationSystem   *ReservationSystem
}

// Функция создания нового фасада для бронирования номера
func NewHotelBookingFacade() *HotelBookingFacade {
	return &HotelBookingFacade{
		availabilityChecker: &AvailabilityChecker{},
		roomSelector:        &RoomSelector{},
		reservationSystem:   &ReservationSystem{},
	}
}

// Функция бронирования номера с использованием фасада
func (h *HotelBookingFacade) bookRoom() (int, error) {
	if !h.availabilityChecker.checkAvailability() {
		return 0, errors.New("Нет свободных номеров")
	}
	room := h.roomSelector.selectRoom()
	h.reservationSystem.confirmReservation()
	return room, nil
}

// func main() {
// 	// Создаём фасад
// 	bookingFacade := NewHotelBookingFacade()
// 	// Пытаемся забронировать номер
// 	room, err := bookingFacade.bookRoom()
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("Бронирование подтверждено, ваш номер:%d", room)
// 	}
// }

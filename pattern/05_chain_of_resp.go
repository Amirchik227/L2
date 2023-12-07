package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
Паттерн "Цепочка вызовов" (Chain of Responsibility) — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков. Каждый обработчик решает, может ли он обработать запрос, и передает запрос следующему обработчику в цепочке.

	Применимость:
	1) Когда есть несколько объектов, способных обработать запрос, и порядок их обработки не фиксирован.
	2) Когда вы хотите передать запрос нескольким объектам, но точно не знаете, какому объекту он будет передан.

	Плюсы:
	1) Уменьшение зависимостей: Отправитель запроса не знает, какой объект в конечном итоге обработает запрос, что уменьшает зависимости между отправителем и получателем.
	2) Гибкость в добавлении новых обработчиков: Легко добавлять новые обработчики в цепочку без изменения существующего кода.
	3) Поддержка отмены операции: Возможность прерывать обработку запроса в середине цепочки.

	Минусы:
	1) Гарантированное выполнение: Не гарантируется, что запрос будет обработан одним из обработчиков. Он может пройти по всей цепочке, и его не обработает никто.
	2) Сложность отладки: Сложно отследить, какой конкретный обработчик обработал запрос, если цепочка довольно сложна.
*/

// Посетитель больницы
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// Интерфейс обработчиков посетителей больницы
type Handler interface {
	execute(*Patient)
	setNext(Handler)
}

// Регигистратура
type Reception struct {
	next Handler
}

// Обработка пациента
func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Регистрация уже произведена")
		r.next.execute(p)
		return
	}
	fmt.Println("Регистрация пациента")
	p.registrationDone = true
	// Перенаправление на следующий обарботчик
	r.next.execute(p)
}

// Установка следующего обработчика цепи
func (r *Reception) setNext(next Handler) {
	r.next = next
}

// Доктор
type Doctor struct {
	next Handler
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Диагноз уже поставлен")
		d.next.execute(p)
		return
	}
	fmt.Println("Доктор ставит диагноз")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Handler) {
	d.next = next
}

// Касса для оплаты медицинских услуг
type Cashier struct {
	next Handler
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Оплата произведена")
	}
	fmt.Println("Кассир принимает оплату клиента")
}

func (c *Cashier) setNext(next Handler) {
	c.next = next
}

// func main() {
// 	// Создаем пациента
// 	patient := &Patient{name: "Фёдор"}
// 	// Создаём регистратуру
// 	reception := &Reception{}
// 	// Создаём доктора
// 	doctor := &Doctor{}
// 	// Создаём кассира
// 	cashier := &Cashier{}

// 	// Получив диагноз от доктора, пациент идёт опалчивать услуги
// 	doctor.setNext(cashier)
// 	// Полученив направление, идет к врачу
// 	reception.setNext(doctor)

// 	reception.execute(patient)
// }

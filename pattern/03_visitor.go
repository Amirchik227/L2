package main

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Паттерн "Посетитель" (Visitor) является поведенческим паттерном проектирования, который позволяет определить новую операцию без изменения классов объектов, над которыми эта операция выполняется. Он позволяет создавать новые операции, не изменяя структуру объектов, над которыми эти операции выполняются.

	Применимость:
	1) Когда необходимо выполнить набор связанных, но различных операций над группой объектов без изменения их классов.
	2) Когда структура объектов изменяется редко, но необходимо добавлять новые операции над этой структурой.

	Плюсы:
	1) Отделение операций от структуры: Паттерн позволяет вынести операции из классов объектов, что упрощает добавление новых операций без изменения существующих классов.
	2) Легкость добавления новых операций: Новые операции могут быть добавлены, создавая новые посетителей, без изменения кода объектов.
	3) Централизованное место для операций: Посетитель предоставляет централизованное место для всех операций, что улучшает поддержку и обеспечивает более чистую структуру кода.

	Минусы:
	1) Сложность добавления новых классов: Добавление новых классов объектов может потребовать изменений в каждом посетителе, что усложняет поддержку системы.
	2) Нарушение инкапсуляции: Посетитель требует открытия интерфейсов объектов для добавления новых операций, что может привести к нарушению инкапсуляции.
*/

//Интерфейс для различных геометрических фигур
type Shape interface {
	getType() string
	accept(Visitor)
}

// Структура для квадрата
type Square struct {
	side int
}

// Данный метод позволяет добавить в программу новые операции,
// не изменяя саму структуру
func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

// Структура для прямоугольника
type Rectangle struct {
	length int
	width  int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "rectangle"
}

// Интерфейс для структур, которе помогут добавить новые операции
// для различных геометрических фигур. Можно создать сколько угодно таких
// структур "посетителей" и не изменяя "посещаемой" структуры
type Visitor interface {
	visitForSquare(*Square)
	visitForRectangle(*Rectangle)
}

// Структура для вычисления площади геометрических фигур
type AreaCalculator struct {
	area int
}

func (c *AreaCalculator) visitForSquare(s *Square) {
	c.area = s.side * s.side
}

func (c *AreaCalculator) visitForRectangle(r *Rectangle) {
	c.area = r.length * r.width
}

func (c *AreaCalculator) getArea() int {
	return c.area
}

// Структура для вычисления периметра геометрических фигур
type PerimeterCalculator struct {
	perimeter int
}

func (c *PerimeterCalculator) visitForSquare(s *Square) {
	c.perimeter = 4 * s.side
}

func (c *PerimeterCalculator) visitForRectangle(r *Rectangle) {
	c.perimeter = 2*r.length + 2*r.width

}

func (c *PerimeterCalculator) getPerimeter() int {
	return c.perimeter
}

// func main() {
// 	// Создадим геометрические фигуры
// 	square := &Square{side: 12}
// 	rectangle := &Rectangle{length: 2, width: 3}

// 	// Создадим структуру для подсчета площади фигур
// 	areaCalculator := &AreaCalculator{}

// 	//Подсчитаем площади фигур и выведем их
// 	square.accept(areaCalculator)
// 	fmt.Println("Площадь квадрата:", areaCalculator.getArea())
// 	rectangle.accept(areaCalculator)
// 	fmt.Println("Площадь прямоугольника:", areaCalculator.getArea())

// 	// Создадим структуру для подсчета периметра фигур
// 	perimeterCalculator := &PerimeterCalculator{}
// 	//Подсчитаем периметры фигур и выведем их
// 	square.accept(perimeterCalculator)
// 	fmt.Println("Периметр квадрата:", perimeterCalculator.getPerimeter())
// 	rectangle.accept(perimeterCalculator)
// 	fmt.Println("Периметр прямоугольника:", perimeterCalculator.getPerimeter())
// }

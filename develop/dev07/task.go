package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// Функция объединяет в себе несколько каналов один
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Канал объединяющий в себе несколько других
	single := make(chan interface{})
	var wg sync.WaitGroup
	// single канал работает пока не закрыт хотя бы один done канал
	wg.Add(len(channels))
	// В отдельной горутине принимаем
	for _, done := range channels {
		// В отдельных горутинах считываем значения из всех каналов и передаем их в single
		go func(c <-chan interface{}) {
			// range будет ждать значения пока не закроется done канал
			for v := range c {
				single <- v
			}
			// После закрытия уменьшаем счетчик
			wg.Done()
		}(done) // Передача канала в функцию
	}
	go func() {
		// ждем пока не будут закрыты все каналы
		wg.Wait()
		// А затем закроем и объединяющий канал
		close(single)
	}()
	return single
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			// для наглядности передадим в канал несколько чисел
			for i := 1; i <= 3; i++ {
				c <- i
			}
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	single := or(
		sig(200*time.Millisecond),
		sig(500*time.Millisecond),
		sig(1*time.Second),
		sig(2*time.Second),
		sig(900*time.Millisecond),
	)
	// Считываем значения из объединяющего канала, пока он не закрыт,
	// или просто ожидаем их если done каналы ничего не передают в single
	for v := range single {
		fmt.Println(v)
	}

	fmt.Printf("Done after %v", time.Since(start))
}

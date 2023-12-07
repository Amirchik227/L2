Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
2
1
4
3
5
6
8
0
0
0
0
...
```
Сначала программа в случайном порядке выведет цифры, переданные в оба вызова функции asChan в случайном порядке. Т.к в asChan после записи в канал мы случайное время ждем, поэтому нельзя сказать точно, в каком порядке будут выведены цифры. Merge будет записывать в канал с цифру из любого канала из которого это будет возможно. После записи всех цифр в каналы a и b функция asChan их закроет, но канал с полученный из фунции merge не будет закрыт. Range будет брать значения из канала, пока он не будет закрыт. Даже если значений в нем нет, он все рано возьмет из него значение по умолчанию (в данном случае 0). Так как закрытие канала с в данной программе не предусмотрено, то брать нули он будет бесконечно.
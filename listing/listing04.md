Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

```
В программе в отдельной горутине присходит запись в небуферизированный канал десяти чисел, которые считваются и выводятся в другой горутине. Однако после этого не происходит закрытия канала. Таким образом, цикл ожидает следующего значения, которое никогда не поступит, поэтому и просходит deadlock.
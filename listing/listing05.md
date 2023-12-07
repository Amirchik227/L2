Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Чтобы переменная интерфейсного типа была равна нил, необходимо, чтобы itab и data были ниловыми. Функция test возвращает указатель на customError, следовательно itab и err не будут ниловым.
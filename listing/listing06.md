Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3 2 3]

```
Слайс состоит из указателя на массив, длины и ёмкости. При передаче в функцию происходит копирование слайса со всеми его внутренностями и указателем на массив в том числе. Получается, что указатели на массив слайсов из main и modifySlice указывают на один и тот же массив, поэтому `i[0] = "3"` повлияет на них обоих. 
Далее происходит добавление в i значения. Если текущей ёмкости хватать не будет, append создаст новый слайс, и в нём будет указатель на новый массив, в который будут скопированы значения из старого, что и происходит в нашем случае. Так как теперь у слайсов разные указатели на разные массивы, то и изменение одного из них не затронет второй.



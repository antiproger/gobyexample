// `for` - это единственный цикл доступный в Go.
// Три стандартных примера использования `for`

package main

import "fmt"

func main() {

	// Стандартный тип с единственным условием
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Классическая инициализация/условие/выражение после `for`
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` без условия будет выполняться до тех пор,
	// пока не выполнится `break` (выход из цикла) или
	// `return`, который завершит функцию с циклом.
	for {
		fmt.Println("loop")
		break
	}

	// Так же Вы можете использовать `continue` для
	// немедленного перехода к следующей итерации цикла
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

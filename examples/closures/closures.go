// Go поддерживает [_анонимные функции_](http://en.wikipedia.org/wiki/Anonymous_function), которые могут образовывать
// <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>замыкания</em></a>. Анонимные функции полезны, когда вы хотите
// определить встроенную функцию, не называя ее.

package main

import "fmt"

// Эта функция intSeq возвращает другую функцию, которую
// мы анонимно определяем в теле intSeq. Возвращенная
// функция связывает переменную i, чтобы
// сформировать замыкание.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Мы вызываем `intSeq`, присваивая результат (функцию)
	// `nextInt`. Это значение функции фиксирует свое
	// собственное значение `i`, которое будет обновляться
	// каждый раз, когда мы вызываем `nextInt`.
	nextInt := intSeq()

	// Посмотрите, что происходит при вызове `nextInt`
	// несколько раз.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Чтобы подтвердить, что состояние является уникальным
	// для этой конкретной функции, создайте и протестируйте
	// новую.
	newInts := intSeq()
	fmt.Println(newInts())
}

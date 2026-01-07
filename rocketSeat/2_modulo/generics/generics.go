package main

import "fmt"

type MyType string

func (MyType) Foo() {
	fmt.Println("Foo method called")
}

func main() {
	foo("Rod")
	foo([]int{1, 2, 3, 4, 5})
	foo2(123)

	var mt MyType = ""
	foo3(mt)

	foo4("42343424")
}

func foo[T any](arg T) {
	fmt.Println(arg)
}
func foo2[T comparable](arg T) { // aceita somente tipos comparáveis
	fmt.Println(arg)
}

type MyConstraint interface {
	Foo()
}

func foo3[T MyConstraint](arg T) {
	arg.Foo()
}

type MyConstraint2 interface {
	int | string | []int
}

func foo4[T MyConstraint2](arg T) {
	fmt.Println(arg)
}

func Contains[T comparable](slice []T, comp T) bool { // T é um tipo genérico que deve ser comparável e isso é garantido pelo constraint comparable
	for _, str := range slice {
		if str == comp {
			return true
		}
	}
	return false
}

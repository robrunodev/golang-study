package main

import (
	"2_modulo/interfaces/bar"
	"fmt"
)

type Animal interface {
	Sound() string
}

type AnimalDog interface {
	Walk() string
}

type Dog struct{}

func (Dog) Sound() string {
	return "Au au Latindo"
}

func (*Dog) Walk() string {
	return "Cachorro andando"
}

func (Dog) InterfaceFoo() {
	fmt.Println("Dog implementando InterfaceFoo")
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

type Cat struct{}

func (d *Cat) Sound() string {
	return "Miau Miau!"
}

func takeAnimal(a Animal) {
	// str, ok := a.(string) // Aqui fazemos uma verificação de tipo durante o runtime o famoso type assertion
	switch t := a.(type) {
	case *Dog:
		{
			t.Sound()
		}
	case *Cat:
		{
			t.Sound()
		}

	}

}

type Pedro struct{}

func (Pedro) String() string {
	return "Olá, eu sou o Pedro!"
}

func main() {
	var a Animal
	dog := Dog{}
	var dogPointer *Dog
	whatDoesThisAnimalSay(dog)
	bar.TakeFoo(dog)

	fmt.Println(a == nil)
	fmt.Println(dogPointer == nil)
	a = dog
	//fmt.Println(a == nil) // false, pois a agora referencia um valor do tipo Dog

	p := Pedro{}
	fmt.Println(p) // Aqui vai printar "Olá, eu sou o Pedro!" pois o fmt.Println chama o método String() automaticamente
}

// Em Go declarar uma interface vazia e um any é exatamente a mesma coisa
// func foo(a interface{}) {
// }
// Ou
// func foo(a any) {
// }

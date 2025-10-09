package main

import (
	"2_modulo/structs_methods/foo"
	"encoding/json"
	"fmt"
)

type MinhaString string

type User struct {
	Name    string
	ID      uint64
	foo.Foo // isso não é uma herança, é uma composição de structs, estamos embedando a struct Foo dentro da struct User
}

func (u User) PrintName() { // isso é a declaração de um método em Go
	fmt.Println(u.Name) // e este método pertence ao tipo User
}

func (u *User) PrintNamePointer(newName string) { // Este método recebe um ponteiro para User e não uma cópia, assim podem ser feitas alterações no objeto original
	u.Name = newName
	fmt.Println(u.Name)
}

func UpdateName(u *User, newName string) { //pointer receiver
	u.Name = newName
}

type UserJSON struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
	Age  int    `json:"idade"` // tags para definir como os campos serão nomeados quando a struct for convertida para JSON
}

func main() {

	user := User{
		"Rafael",
		1, // Lembrando que para declarar sem o nome do campo é necessário estar na ordem correta
		foo.Foo{
			Name: "Rafael Foo",
		},
	}

	user2 := User{
		Name: "Rodrigo",
		ID:   1,
	}
	fmt.Println(user2.Name, user)
	user2.PrintName()
	user.PrintNamePointer("João")
	user.Bar() // Chamando o método Bar que pertence a struct Foo, que está embedada na struct User

	UpdateName(&user, "Maria José") // Passando o endereço(ponteiro) de user para a função UpdateName

	userJson := UserJSON{
		Name: "Rafael",
		ID:   1,
		Age:  30,
	}

	res, err := json.Marshal(userJson)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

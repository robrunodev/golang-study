package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrNotFound = errors.New("not found")

func bar() error {
	//return ErrNotFound
	return SqrtError{"Teste"}
}

func main() {

	err3 := bar()

	var sqrtError SqrtError
	if err3 != nil && errors.As(err3, &sqrtError) { // no errors.As precisa de um ponteiro para fazer uma mutação no valor da variável
		fmt.Println(sqrtError.msg)
		return
	}
	fmt.Println("caiu fora")

	if err3 != nil && errors.Is(err3, ErrNotFound) {
		fmt.Println("Deu o erro not found")
		return
	}
	fmt.Println("Foi pra fora")

	x := -10
	res2, err2 := raizQuadrada(float64(x))

	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(res2)

	a, b := 10, 0

	res, err := dividir(a, b)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	user, err := NewUser(true)

	if err != nil {
		fmt.Println("Algum error na hora de criar o usuário")
		return
	}

	user.Foo()

}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não pode dividir por zero")
	}

	return a / b, nil
}

type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("um erro")
	}
	return &User{}, nil
}

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string {
	return s.msg
}

func raizQuadrada(x float64) (float64, error) {

	if x < 0 {
		return 0, SqrtError{msg: "Não é possível calcular a raiz quadrada de um número negativo"}
	}

	resultado := math.Sqrt(x)
	return resultado, nil
}

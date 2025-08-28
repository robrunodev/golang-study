package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 10000000000
	fmt.Println(numero)

	var numero2 uint64 = 10000000000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 10000
	fmt.Println(numero3)

	// UINT8 = BYTE
	var numero4 byte = 100
	fmt.Println(numero4)

	var numeroReal1 float32 = 124.34
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 14004000000000.34
	fmt.Println(numeroReal2)

	numeroReal3 := 150.50
	fmt.Println(numeroReal3)

	// Strings

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro personalizado")
	fmt.Println(erro)

}

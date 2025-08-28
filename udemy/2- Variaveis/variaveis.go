package main

import "fmt"

func main() {
	var variavel1 = "Variável 1 string"
	variavel2 := "Variável 2 string"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3 string"
		variavel4 string = "Variável 4 string"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5 string", "Variável 6 string"
	fmt.Println(variavel5, variavel6)

	const constante1 = "Constante 1 string"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 // Troca os valores das variáveis
}

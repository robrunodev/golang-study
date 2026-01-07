package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(pacote.Boo)
	// pacote.PrintMinha()

	// fmt.Println(somar(2, 3))

	// a, b := swap(10, 20)
	// fmt.Println(a, b)

	// res, rem := dividir(10, 3)
	// fmt.Println(res, rem)

	// x := somarHighOrder(10)(404)
	// t := somarHighOrder(20)
	// f := t(30)
	// fmt.Println(x) // chamada da função highOrder
	// fmt.Println(f) // chamada da função highOrder

	fmt.Println(somarVariatica(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

// func somar(a, b int) int {
// 	return a + b
// }

// func swap(a, b int) (int, int) {
// 	return b, a
// }

// func dividir(a, b int) (res int, rem int) {
// 	res = a / b
// 	rem = a % b
// 	return res, rem
// }

// high order function
// func somarHighOrder(a int) func(int) int {
// 	return func(b int) int {
// 		return a + b
// 	}
// }

func somarVariatica(nums ...int) int { // Recebe um numero variavel de parametros e é necessário sempre ser o ultimo parametro de uma função
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}

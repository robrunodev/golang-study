package main

import (
	"errors"
	"fmt"
)

func main() {
	// if x := math.Sqrt(4); x < 1 {
	// 	fmt.Println(x)
	// } else if x > 0 {
	// 	fmt.Println("Maior que zero")
	// } else {
	// 	fmt.Println("Menor que zero")
	// }

	meuBool := true
	if meuBool {
		fmt.Println("Verdadeiro")
	}

	if err := doError(); err != nil {
		fmt.Println(err)
	}

}

func doError() error {
	return errors.New("error")
}

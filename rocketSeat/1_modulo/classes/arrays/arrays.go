package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	arr2 := []int{4: 10}
	fmt.Println(arr2)

	const x = 10
	arr3 := [x]string{4: "hello world"}
	fmt.Println(arr3)
}

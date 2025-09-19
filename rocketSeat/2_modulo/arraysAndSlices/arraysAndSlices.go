package main

import "fmt"

func main() {
	// Qual a diferença entre array e slice?
	// O Slice nada mais é do que um "array dinâmico", um array que cresce de tamanho conforme a necessidade.

	arr := [5]int{1, 2, 3, 4, 5} // Array de inteiros com tamanho fixo de 5 elementos
	slice := arr[1:4]
	fmt.Println(arr, slice)

	sliceInitial := []int{1, 2, 3} // Slice de inteiros
	fmt.Println(sliceInitial)

	arrayInts := [5]int{1, 2, 3, 4, 5} // array de inteiros
	sliceInts := arrayInts[:]
	fmt.Println(sliceInts)

	// Todo Slice tem length e capacity
	// length é o tamanho do slice e capacity é a capacidade do slice(quantidade máxima que ele pode crescer)
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := slice3[:0] // length 0, capacity 5
	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(slice4, len(slice4), cap(slice4))

	// Como tudo em Go os slices também tem o valor 0 por padrão
	var slice5 []int                          // slice nil, length 0, capacity 0
	slice6 := []int{}                         // slice vazio, length 0, capacity 0 chamamos de slice literal
	fmt.Println(slice5 == nil, slice6 == nil) // true false porque slice6 é um slice vazio e slice5 é um slice nil

	var filmesNoDB = []string{
		"Star Wars",
		"Matrix",
		"Senhor dos Anéis",
		"Clube da Luta",
		"Pulp Fiction",
		"Forrest Gump",
		"Interestelar",
		"Parasita",
		"Corra!",
		"O Grande Hotel Budapeste",
	}

	// matrix := [][]int{}
	// matrix3D := [][][]int{}

	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filmes := make([]string, 0, len(resultsFromApi)) // length 0, capacity 10

	for _, id := range resultsFromApi {
		filme := filmesNoDB[id-1]
		fmt.Println(len(filmes), cap(filmes))
		filmes = append(filmes, filme)
	}
	fmt.Println(filmes)

	// Full Slice Expression
	arr7 := [5]int{1, 2, 3, 4, 5}
	slice7 := arr7[:2:2]
	fmt.Println(slice7, cap(slice7), len(slice7))

	// Full Slice Expression
	slice8 := []int{1, 2, 3, 4}
	// fmt.Println(slice8[4])
	foo(slice8)

	// slices e arrays são passados por referencia
	fooPointerOrReference(slice8)

}

func foo(slice []int) {
	_ = slice[3] // bound checking é feito em tempo de execução
	// _ = slice[4] // panic: runtime error: index out of range [4] with length 4

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
}

func fooPointerOrReference(slice []int) {
	slice[0] = 123
	fmt.Println(slice)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// Mapas são coleções de chave-valor, onde a chave é única e é usada para acessar o valor correspondente.
	var m map[string]string
	fmt.Println(m == nil)

	// Temos estas duas formas de iniciar um map
	filmes := make(map[string]string, 0) // map vazio
	filmes2 := map[string]string{}       // map vazio literal

	fmt.Println(filmes, filmes2)

	// Um mapa de slices é um mapa onde o valor é um slice.
	// Isso é útil quando queremos armazenar múltiplos valores para uma única chave.
	mSlices := map[string][]int{
		"Rodrigo": {1, 2, 3},
	}
	fmt.Println(mSlices)

	mAddValues := make(map[string]any)
	mAddValues["Nome"] = "Rodrigo"
	mAddValues["Idade"] = 32
	valor, ok := mAddValues["Idade"] // ok é um booleano que indica se a chave existe no mapa
	valor2, ok2 := mAddValues["foo"]
	fmt.Println(valor2, ok2, valor, ok)

	//Deletar uma chave de um mapa
	delete(mAddValues, "Idade")
	valor3, ok3 := mAddValues["Idade"]
	fmt.Println(valor3, ok3)

	// É possível apagar os elementos de um mapa porém ainda manter a sua capacidade
	mTestClear := map[string]any{
		"Nome":  "Rodrigo",
		"Idade": 32,
		"Ativo": true,
	}

	clear(mTestClear)
	fmt.Println(len(mTestClear), mTestClear)

	// Mapas não são concorrentes, ou seja, não são seguros para uso em múltiplas goroutines.
	f := math.NaN()
	f2 := math.NaN()
	nanMap := map[float64]string{
		f:  "Rodrigo",
		f2: "Rodrigo2",
	}

	fmt.Println(nanMap)

	// Como NaN não é igual a NaN, as duas chaves são consideradas diferentes.
	// Não podemos utilizar nem um Delete nem um Update, pois não sabemos qual NaN estamos referenciando.
	valor4, ok4 := nanMap[f]
	fmt.Println(valor4, ok4)

	// A única forma de limpar o mapa é utilizando a função clear
	clear(nanMap)
	fmt.Println(nanMap)

	mIterate := map[string]string{
		"Nome":    "Rodrigo",
		"Idade":   "32",
		"Ativo":   "true",
		"Empresa": "Rocketseat",
	}

	for k, v := range mIterate {
		// Não é recomendado deletar elementos de um mapa enquanto estamos iterando sobre ele,
		// porém em Go isso é possível.
		if k == "Idade" {
			delete(mIterate, k)
		}
		fmt.Println(k, v) // Mapas em Go são desordenados, ou seja, a ordem dos elementos não é garantida.
	}

	teste, okTeste := mIterate["Teste"]
	fmt.Println(teste, okTeste)

}

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um número aleatório será sorteado. Tente acertar. O Número é um inteiro entre 0 e 100")

	x := rand.Int64N(51) // 0 a 50
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("O seu chute precisa ser um inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O Numero sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O Numero sorteado é menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabéns! Você acertou! O número era %d\n"+
					"Você acertou em %d\n"+
					"Essas foram as suas tentativas:%v\n",
				x, i+1, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt

	}

	fmt.Printf(
		"Infelizmente você não acertou o número, que era %d\n"+
			"Você teve 10 tentativas\n"+
			"Essas foram as suas tentativas:%v\n",
		x, chutes,
	)
}

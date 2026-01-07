package main

import (
	"fmt"
	"time"
)

func main() {
	do(1)
	isWeekend(time.Now())
	isWeekend2(time.Now())
	verifyType(10.5)
	verifyType("Hello")
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println("1")
		fallthrough // executa o proximo case mesmo que não seja verdadeiro
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("outro numero")
	}
}

func isWeekend(x time.Time) {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		fmt.Println("Dia de semana")
	default:
		fmt.Println("Fim de semana")
	}
}

func isWeekend2(x time.Time) {
	switch x.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Fim de semana")
	default:
		fmt.Println("Dia de semana")
	}
}

func verifyType(x any) {
	switch t := x.(type) {
	case string:
		takeString(t)
	case int, float32, float64:
		fmt.Println("é do tipo Numérico")
	case time.Time:
		fmt.Println("é do tipo Tempo")
	case nil:
		fmt.Println("é do tipo Localização")
	default:
		fmt.Println("é de outro tipo")
	}
}

func takeString(s string) {
	fmt.Println(s)
}

package main

import "fmt"

func main() {
	x := 10
	p := &x            // Esta variável P é um ponteiro que armazena o endereço de memória da variável X
	fmt.Println(p, *p) // Imprime o endereço de memória armazenado em P e o valor armazenado nesse endereço (valor de X),
	// chamamos isso de desreferenciação

	a := 10
	takeCopy(a)
	fmt.Println(a) // o valor de a continua 10, pois a variável a foi passada por valor, mas caso não queira isso podemos passar por referencia ou por ponteiro

	b := 10
	takePointer(&b) // Aqui nós estamos alterando o endereço de memória de b, por isso o valor retornado é 100
	fmt.Println(b)

	c := create() // A função create retorna um ponteiro para um inteiro
	fmt.Println(c, *c)

	var d *int
	d = nil // nil é o valor zero para ponteiros, ou seja, ele não aponta para lugar nenhum
	takePointer2(d)
	fmt.Println(*d) // Isso vai causar um panic em tempo de execução, pois estamos tentando desreferenciar um ponteiro nulo

}

func takeCopy(i int) {
	i = 100
}

func takePointer(i *int) {
	*i = 100
}

func create() *int {
	x := 10
	return &x
}

func foo(x *int) { // quando o asterisco está na frente do tipo, ele indica que é um ponteiro
	*x = 101 // quando o asterisco está na frente da variável, ele indica desreferenciação
}

func takePointer2(p *int) {
	*p = 20
}

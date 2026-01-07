package pacote

import (
	"fmt"
	foo "myFirstGoProject/pacote/internal"
)

var Boo string = "Hello Boo B"

func PrintMinha() {
	fmt.Println(foo.MinhaString)
}

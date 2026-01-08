package foo

import "fmt"

type Foo struct {
	Name string
}

func (s Foo) Bar() {
	s.Name = "Bar"
	fmt.Println(s.Name)
}

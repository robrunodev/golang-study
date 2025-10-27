package bar

import "2_modulo/interfaces/foo"

type Bar struct {
	Name string
}

func TakeFoo(i foo.Interface) {
	i.InterfaceFoo()
}

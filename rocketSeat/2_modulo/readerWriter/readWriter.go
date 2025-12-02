package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "Hello, World!\n"
	reader := strings.NewReader(str)
	buffer := make([]byte, 2)
	writer := MyWriter{}
	// n, err := reader.Read(buffer)

	// if err != nil {
	// 	panic(err)
	// }

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		writer.Write(buffer[:n])

		// fmt.Println(n)
		// fmt.Println(string(buffer[:n]))
	}

}

type MyWriter struct{}

func (MyWriter) Write(b []byte) (int, error) {
	fmt.Print(string(b))
	return len(b), nil
}

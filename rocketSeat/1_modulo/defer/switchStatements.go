package main

import (
	"database/sql"
	"fmt"
	"os"
)

func main() {
	doDefer()
	openFileDefer()
	filesDefer()
	connectToDB()
}

func doDefer() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Print(1)
}

func openFileDefer() {
	file, err := os.Open("foo.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
}

var arquivos = []string{"a.txt", "b.txt", "c.txt"}

func filesDefer() {
	for _, f := range arquivos {
		func() {
			file, err := os.Open(f)
			if err != nil {
				panic(err)
			}
			defer file.Close()
		}()
	}
}

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("", "")
	if err != nil {

	}
	defer db.Close()
	return db, nil
}

// O Defer segue a ordem Last In First Out
// Ou seja, o ultimo a entrar é o primeiro a sair
// Defer só é executado quando a função em que ele está é retornada

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func goRoutineExample() {

	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	for range n {
		go func() {
			defer wg.Done()

			resp, err := http.Get("https://www.google.com")
			if err != nil {
				panic(err)
			}
			// defer é usado para garantir que o body seja fechado após o uso,
			// sempre precisamos fazer isso para evitar vazamentos de recursos
			defer resp.Body.Close()
			fmt.Println("ok", resp.StatusCode)
		}()

	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

package main

import (
	"sync"

	. "github.com/dyden/go-avanced/05.Singleton/db"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}
	wg.Wait()
}

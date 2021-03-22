package main

import (
	"fmt"
	"sync"
)

var (
	currentBalance int = 0
)

func main() {
	//You can use '--race' for see if the program have some warning of 'DATA CONDITION' ->  'go build --race 01./Sync/main.go' -> './main'
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Desposit(i*100, &wg, &lock)

	}
	wg.Wait()
	fmt.Println(Balance(&lock))

}

/***************************************************************
*                          FUNCTIONS                           *
****************************************************************/
func Desposit(amount int, wg *sync.WaitGroup, locker *sync.RWMutex) {
	defer wg.Done()
	locker.Lock() //LOCK 'balance'
	currentBalance += amount
	locker.Unlock() //UNLOCK 'banalce'
}

func Balance(locker *sync.RWMutex) int {
	locker.RLock() //LOCK 'currentBalance' only for read
	value := currentBalance
	locker.RUnlock() //UNLOCK 'currentBalance' only for read
	return value
}

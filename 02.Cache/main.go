package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	cache := CreateCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%d, %s, %d\n", index, time.Since(start), value)
		}(n)
	}
	wg.Wait()
}

/***************************************************************
*                          STRUCTS                              *
****************************************************************/
type Memory struct {
	f      Function
	cache  map[int]FunctionResult
	locker sync.Mutex
}
type FunctionResult struct {
	value interface{}
	err   error
}

/***************************************************************
*                          FUNCTIONS                           *
****************************************************************/

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Function func(key int) (interface{}, error)

func CreateCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.locker.Lock()
	result, exists := m.cache[key]
	m.locker.Unlock()
	if !exists {
		m.locker.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.locker.Unlock()
	}
	return result.value, result.err
}
func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

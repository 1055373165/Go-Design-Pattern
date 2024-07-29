package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var once sync.Once
var instance *singleton

func GetInstance(i int, wg *sync.WaitGroup) *singleton {
	defer wg.Done()

	once.Do(func() {
		instance = new(singleton)
		fmt.Println("goroutine", i, "initilized instance...")
	})

	return instance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go GetInstance(i, &wg)
	}
	wg.Wait()
}

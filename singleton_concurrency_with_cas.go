package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct {
}

func (s *singleton) doSomething() {
	fmt.Println("singleton do something...")
}

var lock sync.Mutex
var instance *singleton
var initilized uint32

func GetInstance(idx int, wg *sync.WaitGroup) *singleton {
	defer wg.Done()

	if atomic.LoadUint32(&initilized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		atomic.StoreUint32(&initilized, 1)
		instance = new(singleton)
		fmt.Println("goroutine", idx, "initilized singleton...")
	}

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

package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var lock sync.Mutex
var instance *singleton

func GetInstance(idx int, wg *sync.WaitGroup) *singleton {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		fmt.Printf("goroutine%d 创建了单例", idx)
		instance = new(singleton)
	}
	wg.Done()
	return instance
}

func (s *singleton) DoSomething() {
	fmt.Println("单例 do something...")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go GetInstance(i, &wg)
	}
	wg.Wait()
}

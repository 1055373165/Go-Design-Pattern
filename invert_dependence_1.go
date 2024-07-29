package main

import "fmt"

// 抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层（面向抽象层编程，只需要实现 Car 接口）
type Benz struct{}

func (bz *Benz) Run() {
	fmt.Println("Benz is running...")
}

type BMW struct {
}

func (bz *BMW) Run() {
	fmt.Println("BMW is running...")
}

// 面向抽象层编程，只需要实现 Driver 接口
type ZhangSan struct{}

func (zs ZhangSan) Drive(car Car) {
	fmt.Println("张三开汽车...")
	car.Run()
}

type LiSi struct{}

func (ls LiSi) Drive(car Car) {
	fmt.Println("李四开汽车...")
	car.Run()
}

// 业务逻辑层
func main() {
	// zhangsan ---> benz
	var benz Car
	benz = new(Benz)
	var z3 Driver
	z3 = new(ZhangSan)

	z3.Drive(benz)

	// lisi ---> bmw
	var bmw Car
	bmw = new(BMW)
	var lisi Driver
	lisi = new(LiSi)

	lisi.Drive(bmw)
}

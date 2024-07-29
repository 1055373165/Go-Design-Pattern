package main

import "fmt"

// 抽象层
type Fruit interface {
	Show()
}

// 实现层（面向抽象编程）
type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("我是苹果...")
}

type Banana struct {
}

func (b *Banana) Show() {
	fmt.Println("我是香蕉...")
}

type Pear struct {
}

func (a *Pear) Show() {
	fmt.Println("我是梨子...")
}

// 工厂模块
type Factory struct {
}

func (f *Factory) CraeteFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple) // 满足多态条件的赋值，父类指针指向子类对象
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}

	return fruit
}

func main() {
	factory := new(Factory)
	apple := factory.CraeteFruit("apple")
	apple.Show()

	banana := factory.CraeteFruit("banana")
	banana.Show()

	pear := factory.CraeteFruit("pear")
	pear.Show()
}

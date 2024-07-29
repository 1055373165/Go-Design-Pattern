package main

import "fmt"

// 抽象层
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

// 实现层（面向接口编程）
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

type AppleFactory struct {
}

func (af *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

type BananaFactory struct {
}

func (bf *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)
	return fruit
}

type PearFactory struct {
}

func (pf *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Pear)
	return fruit
}

// 业务逻辑层
func main() {
	var appleFactory AbstractFactory
	appleFactory = new(AppleFactory)
	var apple Fruit
	apple = appleFactory.CreateFruit()
	apple.Show()

	var bananaFactory AbstractFactory
	var banana Fruit
	bananaFactory = new(BananaFactory)
	banana = bananaFactory.CreateFruit()
	banana.Show()

	pearFactory := new(PearFactory)
	pear := pearFactory.CreateFruit()
	pear.Show()

	hamiMelonFactory := new(HamiMelonFactory)
	hamiMelon := hamiMelonFactory.CreateFruit()
	hamiMelon.Show()
}

type HamiMelon struct {
}

func (hm *HamiMelon) Show() {
	fmt.Println("我是哈密瓜...")
}

type HamiMelonFactory struct {
}

func (hf *HamiMelonFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(HamiMelon)
	return fruit
}

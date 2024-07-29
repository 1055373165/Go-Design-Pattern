package main

import "fmt"

type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()
	WantAddThings() bool
}

type template struct {
	b Beverage
}

// 封装固定的模版，制作饮料
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	// 固定流程
	t.b.BoilWater() // 实际执行的是子类具体的实现方法
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

type MakeCoffee struct {
	template // 继承模版
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	// b 是 Beverage，是 MakeCoffee 接口，需要给接口赋值让其指向具体的子类
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水煮到 100 摄氏度")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("用热水冲咖啡")
}

func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将冲好的咖啡倒入陶瓷杯中")
}

func (mc *MakeCoffee) AddThings() {
	fmt.Println("再往陶瓷杯中添加糖和牛奶")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

type MakeTee struct {
	template
}

func NewMakeTee() *MakeTee {
	makeTee := new(MakeTee)
	makeTee.b = makeTee
	return makeTee
}

func (mt *MakeTee) BoilWater() {
	fmt.Println("将水煮到 80 摄氏度")
}

func (mt *MakeTee) Brew() {
	fmt.Println("用热水冲茶叶")
}

func (mt *MakeTee) PourInCup() {
	fmt.Println("将冲好的茶水倒入茶壶中")
}

func (mt *MakeTee) AddThings() {
	fmt.Println("再往茶壶中添加柠檬")
}

func (mc *MakeTee) WantAddThings() bool {
	return false
}

func main() {
	// 制作一杯咖啡
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage()
	fmt.Println("=========================")
	// 制作一杯茶
	makeTee := NewMakeTee()
	makeTee.MakeBeverage()
}

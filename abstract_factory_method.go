package main

import "fmt"

// 抽象层
type FruitFactory interface {
	CreateApple() Apple
	CreateBanana() Banana
	CreatePear() Pear
}

type Apple interface {
	ShowApple()
}

type Banana interface {
	ShowBanana()
}

type Pear interface {
	ShowPear()
}

// 实现层（面向抽象编程）
type ChinaApple struct {
}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("我是中国苹果！")
}

type ChinaBanana struct {
}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("我是中国香蕉！")
}

type ChinaPear struct {
}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("我是中国梨子！")
}

type AmericanApple struct {
}

func (aa *AmericanApple) ShowApple() {
	fmt.Println("我是美国香蕉")
}

type AmericanBanana struct {
}

func (ab *AmericanBanana) ShowBanana() {
	fmt.Println("我是美国香蕉")
}

type AmericanPear struct {
}

func (ap *AmericanPear) ShowPear() {
	fmt.Println("我是美国梨子")
}

type RussiaApple struct {
}

func (ra *RussiaApple) ShowApple() {
	fmt.Println("我是俄罗斯苹果")
}

type RussiaBanana struct {
}

func (rb *RussiaBanana) ShowBanana() {
	fmt.Println("我是俄罗斯香蕉")
}

type RussianPear struct {
}

func (rp *RussianPear) ShowPear() {
	fmt.Println("我是俄罗斯梨子")
}

type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateApple() Apple {
	var apple Apple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() Banana {
	var banana Banana
	banana = new(ChinaBanana)
	return banana
}

func (cf *ChinaFactory) CreatePear() Pear {
	var pear Pear
	pear = new(ChinaPear)
	return pear
}

type AmericanFactory struct {
}

func (af *AmericanFactory) CreateApple() Apple {
	var apple Apple
	apple = new(AmericanApple)
	return apple
}

func (af *AmericanFactory) CreateBanana() Banana {
	var banana Banana
	banana = new(AmericanBanana)
	return banana
}

func (af *AmericanFactory) CreatePear() Pear {
	var pear Pear
	pear = new(AmericanPear)
	return pear
}

type RussiaFactory struct {
}

func (af *RussiaFactory) CreateApple() Apple {
	var apple Apple
	apple = new(RussiaApple)
	return apple
}

func (af *RussiaFactory) CreateBanana() Banana {
	var banana Banana
	banana = new(RussiaBanana)
	return banana
}

func (af *RussiaFactory) CreatePear() Pear {
	var pear Pear
	pear = new(RussianPear)
	return pear
}

func main() {
	// 中国砀山梨
	var factory FruitFactory
	var pear Pear
	factory = new(ChinaFactory)
	pear = factory.CreatePear()
	pear.ShowPear()
	// 美国香蕉
	var banana Banana
	factory = new(AmericanFactory)
	banana = factory.CreateBanana()
	banana.ShowBanana()
	// 俄罗斯苹果
	var apple Apple
	factory = new(RussiaFactory)
	apple = factory.CreateApple()
	apple.ShowApple()
}

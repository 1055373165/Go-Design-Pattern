package main

import "fmt"

type SubSystemA struct {
}

func (sa *SubSystemA) MethodA() {
	fmt.Println("子系统方法 A")
}

type SubSystemB struct {
}

func (sb *SubSystemB) MethodB() {
	fmt.Println("子系统方法 B")
}

type SubSystemC struct {
}

func (sc *SubSystemC) MethodC() {
	fmt.Println("子系统方法 C")
}

type SubSystemD struct {
}

func (sc *SubSystemD) MethodD() {
	fmt.Println("子系统方法 D")
}

// 外观模式，提供一个外观类，简化成一个简单的接口供使用
type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

// 包裹方法
func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *Facade) MethodTwo() {
	f.c.MethodC()
	f.d.MethodD()
}

func (f *Facade) MethodThree() {
	f.a.MethodA()
	f.b.MethodB()
	f.c.MethodC()
	f.d.MethodD()
}

func main() {
	// 如果不适用外观模式，调用 A 和 B、C 和 D 如下所示
	sa := new(SubSystemA)
	sa.MethodA()
	sb := new(SubSystemB)
	sb.MethodB()
	sc := new(SubSystemC)
	sc.MethodC()
	sd := new(SubSystemD)
	sd.MethodD()

	fmt.Println("=========================")

	// 使用外观模式分别调用 A 和 B 的方法以及分别调用四个子系统方法
	f := Facade{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
		d: new(SubSystemD),
	}
	f.MethodOne()
	fmt.Println("=========================")
	f.MethodThree()
}

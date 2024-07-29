package main

import "fmt"

// 适配的目标
type V5 interface {
	Use5V()
}

// 被适配的橘色，适配者
type V220 struct {
}

func (v *V220) Use220V() {
	fmt.Println("连接 220V 电压充电")
}

type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone 进行了充电...")
	p.v.Use5V()
}

// 适配器类
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	a.v220.Use220V()
	fmt.Println("使用适配器调节电压 220V 为 5V 进行充电...")
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{
		v220: v220,
	}
}

// ========= 业务逻辑层 =========
func main() {
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()
}

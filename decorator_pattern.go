package main

import "fmt"

// ========= 抽象层 =========
type Phone interface {
	Show()
}

// 抽象的装饰器，装饰器的基础类（该类本应是 interface，但是 go 的 interface 不可以有成员属性）
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

// ========= 实现层 =========
type Huawei struct {
}

func (hw *Huawei) Show() {
	fmt.Println("拿出华为手机...")
}

type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("拿出苹果手机...")
}

// 主要继承 Phone 的成员属性
type StickFilmDecorate struct {
	Decorator
}

func (sfd *StickFilmDecorate) Show() {
	sfd.phone.Show()
	fmt.Println("现在是有了贴膜的手机")
}

func NewStickFilmDecorate(phone Phone) Phone {
	return &StickFilmDecorate{Decorator{phone: phone}}
}

type PhoneCaseDecorate struct {
	Decorator
}

func (pcd *PhoneCaseDecorate) Show() {
	pcd.phone.Show()
	fmt.Println("现在是有了手机壳的手机")
}

func NewPhoneCaseDecorate(phone Phone) Phone {
	return &PhoneCaseDecorate{Decorator{phone: phone}}
}

// ========= 业务逻辑层 =========
func main() {
	var huawei, apple Phone
	huawei, apple = new(Huawei), new(Apple)
	huawei.Show()

	fmt.Println("============================")

	var stickFilmHuawei Phone
	stickFilmHuawei = NewStickFilmDecorate(huawei)
	stickFilmHuawei.Show()

	fmt.Println("============================")

	var phoneCaseApple Phone
	phoneCaseApple = NewPhoneCaseDecorate(apple)
	phoneCaseApple.Show()

	fmt.Println("============================")

	var stickCasePhone Phone
	stickCasePhone = NewPhoneCaseDecorate(stickFilmHuawei)
	// Call Chain: PhoneCaseDecorate.Show -> StickFilmDecorate.Show -> huawei.Show()
	// result in turn
	stickCasePhone.Show()
}

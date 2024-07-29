package main

import "fmt"

type Attack interface {
	Flight()
}

type Hero struct {
	Name   string
	attack Attack // 攻击方式
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能")
	h.attack.Flight()
}

// 属于攻击方式的一种s
type Dabaojian struct {
}

func (d *Dabaojian) Flight() {
	fmt.Println("使用了大宝剑技能，斩杀 33% 的血量。")
}

type PowerOff struct {
}

func (p *PowerOff) Shutdown() {
	fmt.Println("因为释放了大招，电脑支撑不住直接关机")
}

type Adapter struct {
	p *PowerOff
}

func (a *Adapter) Flight() {
	a.p.Shutdown()
}

func NewAdapter(p *PowerOff) *Adapter {
	return &Adapter{
		p: p,
	}
}

func main() {
	gailun := Hero{Name: "盖伦", attack: new(Adapter)}
	gailun.Skill()
}

package main

import "fmt"

type Television struct{}

func (t *Television) On() {
	fmt.Println("打开电视机")
}

func (t *Television) Off() {
	fmt.Println("关闭电视机")
}

type GameConsoles struct{}

func (gc *GameConsoles) On() {
	fmt.Println("打开游戏机")
}

func (gc *GameConsoles) Off() {
	fmt.Println("关闭游戏机")
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("开灯")
}

func (l *Light) Off() {
	fmt.Println("关灯")
}

type Microphone struct{}

func (m *Microphone) On() {
	fmt.Println("打开麦克风")
}

func (m *Microphone) Off() {
	fmt.Println("关闭麦克风")
}

type Audio struct{}

func (a *Audio) On() {
	fmt.Println("打开音响")
}

func (a *Audio) Off() {
	fmt.Println("关闭音响")
}

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("打开投影仪")
}

func (p *Projector) Off() {
	fmt.Println("关闭投影仪")
}

// 家庭影院（外观模式）
type HouseholdPlayerFacade struct {
	light Light
	tv    Television
	xbox  GameConsoles
	mp    Microphone
	proj  Projector
	audio Audio
}

func (hpf *HouseholdPlayerFacade) StartMovieMode() {
	fmt.Println("一键开启看电影模式")
	hpf.light.On()
	hpf.tv.On()
	hpf.audio.On()
	hpf.proj.On()
	hpf.mp.Off()
	hpf.xbox.Off()
}

func (hpf *HouseholdPlayerFacade) StartKTVMode() {
	fmt.Println("一键开启KTV模式")
	hpf.proj.Off()
	hpf.xbox.Off()
	hpf.light.Off()
	hpf.tv.On()
	hpf.mp.On()
	hpf.audio.On()
}

func (hpf *HouseholdPlayerFacade) StartGameMode() {
	fmt.Println("一键开启游戏模式")
	hpf.tv.Off()
	hpf.proj.Off()
	hpf.mp.Off()
	hpf.xbox.On()
	hpf.light.On()
	hpf.audio.On()
}

func main() {
	// 不使用外观模式开启看电影模式
	var l = new(Light)
	var t = new(Television)
	var proj = new(Projector)
	var a = new(Audio)
	l.On()
	t.On()
	proj.On()
	a.On()

	fmt.Println("=========================")
	// 使用外观模式开启看电影模式
	var facade *HouseholdPlayerFacade
	facade = new(HouseholdPlayerFacade)
	facade.StartMovieMode()
	fmt.Println("=========================")
	// 使用外观模式开启 KTV 模式
	facade.StartKTVMode()
	fmt.Println("=========================")
	// 使用外观模式开启游戏模式
	facade.StartGameMode()
}

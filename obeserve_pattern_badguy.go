package main

import "fmt"

// 监听老师是否在来教师的路上
type Listener interface {
	OnTeacherComing()
}

type Notifier interface {
	AddListener(Listener)
	RemoveListener(Listener)
	Notify()
}

type ZhangSan struct {
	inProgress string
}

func (zs *ZhangSan) OnTeacherComing() {
	fmt.Printf("老师来了！张三停止%s\n", zs.inProgress)
}

type LiSi struct {
	inProgress string
}

func (ls *LiSi) OnTeacherComing() {
	fmt.Printf("老师来了！李四停止%s\n", ls.inProgress)
}

type WangWu struct {
	inProgress string
}

func (ww *WangWu) OnTeacherComing() {
	fmt.Printf("老师来了！王五停止%s\n", ww.inProgress)
}

func (zs *ZhangSan) do() {
	fmt.Printf("老师没来，张三放肆的%s\n", zs.inProgress)
}

func (ls *LiSi) do() {
	fmt.Printf("老师没来，李四放肆的%s\n", ls.inProgress)
}

func (ww *WangWu) do() {
	fmt.Printf("老师没来，王五放肆的%s\n", ww.inProgress)
}

// 班长观察老师的动向，因为他被其中几个同学收买了
type ClassMonitor struct {
	listener []Listener // 金主列表
}

func (cm *ClassMonitor) AddListener(l Listener) {
	cm.listener = append(cm.listener, l)
}

func (cm *ClassMonitor) RemoveListener(removing Listener) {
	for idx, l := range cm.listener {
		if l == removing {
			cm.listener = append(cm.listener[:idx], cm.listener[idx+1:]...)
			break
		}
	}
}

func (cm *ClassMonitor) Notify() {
	for _, l := range cm.listener {
		l.OnTeacherComing()
	}
}

func main() {
	stu1 := &ZhangSan{
		inProgress: "抄作业",
	}
	stu2 := &LiSi{
		inProgress: "玩王者荣耀",
	}
	stu3 := &WangWu{
		inProgress: "看李四玩王者荣耀",
	}
	stu1.do()
	stu2.do()
	stu3.do()

	monitor := new(ClassMonitor)
	monitor.AddListener(stu1)
	monitor.AddListener(stu2)
	monitor.AddListener(stu3)

	fmt.Println("=========================")

	monitor.Notify()
}

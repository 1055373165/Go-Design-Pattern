package main

import "fmt"

// ========= 抽象层 =========
type BeautyWoman interface {
	MakeEyesWithMan()
	HappyWithMan()
}

// ========= 实现层 =========
type PanJinLian struct {
}

func (pjl PanJinLian) MakeEyesWithMan() {
	fmt.Println("潘金莲对西门抛了一个媚眼...")
}

func (pjl PanJinLian) HappyWithMan() {
	fmt.Println("潘金莲和西门约会..")
}

type WangPo struct {
	woman BeautyWoman
}

func NewProxy(w BeautyWoman) WangPo {
	return WangPo{
		woman: w,
	}
}

func (wp WangPo) MakeEyesWithMan() {
	wp.woman.MakeEyesWithMan()
}

func (wp WangPo) HappyWithMan() {
	wp.woman.HappyWithMan()
}

// ========= 业务逻辑层 =========
// 西门
func main() {
	// 大官人想要找潘金莲，让王婆来安排
	wangpo := NewProxy(new(PanJinLian))
	// 王婆命令金莲抛媚眼
	wangpo.MakeEyesWithMan()
	// 王婆命令金莲和西门约会
	wangpo.HappyWithMan()
}

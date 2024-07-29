package main

import "fmt"

type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("治疗鼻子")
}

type List interface {
	Treat()
}

type ListForTreatEye struct {
	d *Doctor
}

func (l *ListForTreatEye) Treat() {
	l.d.treatEye()
}

type ListForTreatNose struct {
	d *Doctor
}

func (l *ListForTreatNose) Treat() {
	l.d.treatNose()
}

// 护士---调用命令者
type Nurse struct {
	Lists []List // 收集病单
}

func (n *Nurse) Notify() {
	if n.Lists == nil {
		return
	}

	for _, list := range n.Lists {
		list.Treat() // 执行病单已经绑定的命令（这里会调用病单已经绑定的医生的诊断方法）
	}
}

func main() {
	// 依赖病单，通过填写病单，让医生看病
	doctor := new(Doctor)

	// 治疗眼睛和鼻子的病单
	listEye := ListForTreatEye{doctor}
	listNose := ListForTreatNose{doctor}

	// 护士收集病单
	nurse := new(Nurse)
	nurse.Lists = append(nurse.Lists, &listEye)
	nurse.Lists = append(nurse.Lists, &listNose)

	// 执行病单指令
	nurse.Notify()
}

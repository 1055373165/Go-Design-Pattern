package main

import "fmt"

type Doctor struct {
}

func (d *Doctor) eyeTreat() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) noseTreat() {
	fmt.Println("医生治疗鼻子")
}

func main() {
	doctor := new(Doctor)
	doctor.eyeTreat()
	doctor.noseTreat()
}

package main

import "fmt"

type Benz struct{}

func (b Benz) Run() {
	fmt.Println("Benz is running...")
}

type BMW struct{}

func (b BMW) Run() {
	fmt.Println("BMW is running...")
}

type ZhangSan struct {
}

func (zs ZhangSan) DriveBenz() {
	fmt.Println("zhangsan drive benz")
}

func (zs ZhangSan) DriveBMW() {
	fmt.Println("zhangsan drive bmw")
}

type LiSi struct {
}

func (ls *LiSi) DriveBenz() {
	fmt.Println("lisi drive benz")
}

func (ls *LiSi) DriveBMW() {
	fmt.Println("lisi drive bmw")
}

func main() {
	var zs ZhangSan
	var ls LiSi

	// zhangsan drive bmw
	zs.DriveBMW()
	// lisi drive benz
	ls.DriveBenz()
}

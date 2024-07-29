package main

import "fmt"

type ClothesShop struct{}

func (cs *ClothesShop) Wear() {
	fmt.Println("逛街装扮...")
}

type ClothesWork struct{}

func (cw *ClothesWork) Wear() {
	fmt.Println("工作装扮...")
}

func main() {
	var cs ClothesShop
	var cw ClothesWork

	// 逛街场景
	cs.Wear()
	// 工作场景
	cw.Wear()
}

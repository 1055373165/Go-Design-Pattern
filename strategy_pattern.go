package main

import "fmt"

type WeaponStrategy interface {
	Attack()
}

type AK47 struct {
}

func (ak *AK47) Attack() {
	fmt.Println("使用 ak47 发起攻击...")
}

type UZI struct {
}

func (uzi *UZI) Attack() {
	fmt.Println("使用 uzi 发起攻击...")
}

type Hero struct {
	strategy WeaponStrategy
}

func (h *Hero) Fight() {
	h.strategy.Attack()
}

func (h *Hero) SetWeaponStrategy(w WeaponStrategy) {
	h.strategy = w
}

func main() {
	h := Hero{}

	h.SetWeaponStrategy(new(AK47))
	h.Fight()

	h.SetWeaponStrategy(new(UZI))
	h.Fight()
}

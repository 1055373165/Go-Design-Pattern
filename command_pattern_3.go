package main

import "fmt"

type BBQlMaster struct {
}

func (b *BBQlMaster) BBQChicken() {
	fmt.Println("开始烤鸡腿...")
}

func (b *BBQlMaster) BBQMutton() {
	fmt.Println("开始烤羊肉...")
}

type BBQOrder interface {
	PlaceOrder()
}

type ChickenOrder struct {
	master *BBQlMaster
}

func (co *ChickenOrder) PlaceOrder() {
	fmt.Println("点个鸡腿...")
	co.master.BBQChicken()
}

type MuttonOrder struct {
	master *BBQlMaster
}

func (mo *MuttonOrder) PlaceOrder() {
	fmt.Println("点个羊肉...")
	mo.master.BBQMutton()
}

type Waitress struct {
	orderList []BBQOrder
}

func (w *Waitress) Notify() {
	if w.orderList == nil {
		return
	}

	for _, order := range w.orderList {
		order.PlaceOrder()
	}
}

func main() {
	// 真正的大师永远怀着一颗学徒的心
	master := new(BBQlMaster)
	// 具体烧烤订单
	chickenOrder := ChickenOrder{master: master}
	muttonOrder := MuttonOrder{master: master}

	// 服务员收集订单
	var w Waitress
	w.orderList = append(w.orderList, &chickenOrder)
	w.orderList = append(w.orderList, &muttonOrder)

	// 由服务员将收集到的订单批量通知烧烤师
	w.Notify()
}

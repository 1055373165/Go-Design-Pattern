package main

import "fmt"

type Good struct {
	Kind string // 商品种类
	Fact bool   // 商品的真伪
}

// ========= 抽象层 =========
type Shopping interface {
	Buy(Good)
}

// ========= 实现层 =========
type AmericanShopping struct {
}

func (as *AmericanShopping) Buy(good Good) {
	fmt.Printf("在美国买了一件商品, 商品种类是: %s, 商品真伪为真: %v\n", good.Kind, good.Fact)
}

type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(good Good) {
	fmt.Printf("在韩国买了一件商品, 商品种类是: %s, 商品真伪为真: %v\n", good.Kind, good.Fact)
}

type AfricaShopping struct {
}

func (afs *AfricaShopping) Buy(good Good) {
	fmt.Printf("在非洲买了一件商品, 商品种类是: %s, 商品真伪为真: %v\n", good.Kind, good.Fact)
}

type ShoppingProxy struct {
	shopping Shopping
}

func NewProxy(shopping Shopping) Shopping {
	return &ShoppingProxy{
		shopping: shopping,
	}
}

func (sp *ShoppingProxy) IdentityAuthenticity(good Good) bool {
	if !good.Fact {
		return false
	}
	return true
}

func (sp *ShoppingProxy) PassSecurityCheck() {
	fmt.Println("海外代购通过国际快运寄件, 货物通过海关安检...")
}

func (sp *ShoppingProxy) Buy(good Good) {
	if sp.IdentityAuthenticity(good) {
		fmt.Println("海外代购验货, 货品是真品, 继续购买!")
		sp.shopping.Buy(good)
		sp.PassSecurityCheck()
		return
	}
	fmt.Println("海外代购验货, 货品是仿品, 放弃购买!")
}

// ========= 业务逻辑层 =========
func main() {
	goods := []Good{
		{
			Kind: "计算机",
			Fact: true,
		},
		{
			Kind: "化妆品",
			Fact: false,
		},
		{
			Kind: "咖啡",
			Fact: true,
		},
	}

	var as AmericanShopping
	as.Buy(goods[0])

	var ks KoreaShopping
	ks.Buy(goods[1])

	var afs AfricaShopping
	afs.Buy(goods[2])

	fmt.Println("===============美国代购===============")
	var sp Shopping
	sp = NewProxy(&as)
	sp.Buy(goods[0])

	fmt.Println("===============韩国代购===============")
	sp = NewProxy(&ks)
	sp.Buy(goods[1])

	fmt.Println("===============非洲代购===============")
	sp = NewProxy(&afs)
	sp.Buy(goods[2])
}

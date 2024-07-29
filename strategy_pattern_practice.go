package main

import "fmt"

type SellStrategy interface {
	GetPrice(float64) float64
}

type StrategyA struct {
}

func (a *StrategyA) GetPrice(price float64) float64 {
	if price < 0 {
		fmt.Println("价格出现严重错误")
		return 0
	}

	fmt.Println("打折促销策略 1 生效中，总价打八折")
	return 0.8 * price
}

type StrategyB struct {
}

func (b *StrategyB) GetPrice(price float64) float64 {
	if price < 0 {
		fmt.Println("价格出现严重错误")
		return 0
	}

	if price >= 200 {
		fmt.Println("该商品使用打折促销策略 2 ，满 200 减 100")
		return price - 100
	}

	fmt.Println("未满足打折促销策略 2 的条件，不参与打折")
	return price
}

type Goods struct {
	name     string
	price    float64
	strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy) {
	g.strategy = s
}

// 模拟收银
func (g *Goods) SellPrice() float64 {
	fmt.Printf("商品原价为: %f\n", g.price)
	return g.strategy.GetPrice(g.price)
}

func main() {
	var GoodsList = []Goods{
		{
			name:     "榴莲",
			price:    222.222,
			strategy: new(StrategyB),
		},
		{
			name:     "黄瓜",
			price:    5.20,
			strategy: new(StrategyA),
		},
		{
			name:     "大枣",
			price:    20.00,
			strategy: new(StrategyA),
		},
		{
			name:     "xxx保健品",
			price:    2888,
			strategy: new(StrategyB),
		},
	}

	for _, goods := range GoodsList {
		fmt.Printf("打折后的商品总价为: %.2f\n", goods.SellPrice())
		fmt.Println()
	}
}

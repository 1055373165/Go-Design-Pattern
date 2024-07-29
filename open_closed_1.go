package main

import "fmt"

type Banker interface {
	DoBusiness()
}

func BankBusiness(banker Banker) {
	// 通过接口向下来调用（实现多态 ）
	banker.DoBusiness()
}

// banker in charge of deposits
type DepositBanker struct {
}

func (db DepositBanker) DoBusiness() {
	fmt.Println("存款业务员办理了存款业务...")
}

type TransferBanker struct {
}

func (tb TransferBanker) DoBusiness() {
	fmt.Println("转账业务员办理了转账业务...")
}

type ShareBanker struct {
}

func (sb ShareBanker) DoBusiness() {
	fmt.Println("股票业务员办理了股票业务...")
}

func main() {
	var db DepositBanker
	var tb TransferBanker
	var sb ShareBanker

	// deposit business
	BankBusiness(&db)

	// transfer business
	BankBusiness(&tb)

	// share business
	BankBusiness(&sb)
}

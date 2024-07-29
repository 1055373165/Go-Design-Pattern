package main

import "fmt"

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
	db.DoBusiness()

	// transfer business
	tb.DoBusiness()

	// share business
	sb.DoBusiness()
}

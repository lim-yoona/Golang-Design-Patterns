package main

import "Golang-Design-Patterns/principle"

// main.go用于测试代码
func main() {
	sb := principle.SaveBank{}
	tb := principle.TranBank{}
	gb := principle.GetMoneyBank{}
	principle.BankBusiness(&sb)
	principle.BankBusiness(&tb)
	principle.BankBusiness(&gb)
}

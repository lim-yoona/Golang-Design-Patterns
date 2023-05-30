package main

import "Golang-Design-Patterns/principle"

// main.go用于测试代码
func main() {
	// 测试principle/开闭原则.go中的多态
	sb := principle.SaveBank{}
	tb := principle.TranBank{}
	gb := principle.GetMoneyBank{}
	principle.BankBusiness(&sb)
	principle.BankBusiness(&tb)
	principle.BankBusiness(&gb)
}

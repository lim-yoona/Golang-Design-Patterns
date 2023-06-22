package main

import (
	"Golang-Design-Patterns/creational"
	"Golang-Design-Patterns/principle"
)

// main.go用于测试代码
func main() {
	// 测试principle/开闭原则.go中的多态
	sb := principle.SaveBank{}
	tb := principle.TranBank{}
	gb := principle.GetMoneyBank{}
	principle.BankBusiness(&sb)
	principle.BankBusiness(&tb)
	principle.BankBusiness(&gb)

	// 测试principle/依赖倒转原则.go
	principle.Business()

	// 测试principle/合成复用原则.go
	principle.TestCat()

	// 简单工厂模式
	creational.Business()
}

// 缺点：简单工厂模式违反了开闭原则，新增加一个水果的时候需要修改工厂模块
// 缺点：工厂类职责过重，一旦不能工作，系统收到影响

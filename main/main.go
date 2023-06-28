package main

import (
	"Golang-Design-Patterns/behavioral"
	"Golang-Design-Patterns/creational"
	"Golang-Design-Patterns/principle"
	"fmt"
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

	// 工厂方法模式
	creational.FacMethod()

	// 抽象工厂方法模式
	creational.AbstractFacPattern()

	// 抽象工厂方法模式New Case
	creational.CumputerTest()

	// "饿汉式"单例模式
	sc1 := creational.GetInstance()
	sc1.Output()
	sc2 := creational.GetInstance()
	if sc1 == sc2 {
		fmt.Println("sc1==sc2")
	}
	sc1 = nil
	if sc1 == sc2 {
		fmt.Println("sc1==sc2")
	} else {
		fmt.Println("sc1!=sc2")
	}
	fmt.Println("sc1=", sc1)
	fmt.Println("sc2=", sc2)
	sc3 := creational.GetInstance()
	fmt.Println("sc3=", *sc3)

	// 测试懒汉式并发安全单例模式
	ins1 := creational.GetInstance1()
	ins1.Output1()
	ins2 := creational.GetInstance1()
	if ins1 == ins2 {
		fmt.Println("ins1 == ins2")
	}

	// 代理模式
	behavioral.ProxyBuy()
}

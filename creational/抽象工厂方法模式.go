package creational

import "fmt"

// 工厂方法模式，每次生产一个产品的时候，都要有对应的一个工厂
// 为了降低工厂数量，提出了抽象工厂方法模式
// 抽象工厂方法模式的使用场景不是很多

// 抽象工厂模式种，将产品分为“产品族”与“产品等级结构”
// “产品族”是包含所有产品种类的某一个系列（例如：中国水果（其中包括中国苹果、中国香蕉、中国梨等）、日本水果、美国水果）
// “产品等级结构”是某一个产品的不同型号（例如：苹果（其中包括中国苹果、日本苹果、美国苹果）、香蕉、梨）
// “产品等级结构“应该是相对来说比较稳定的，很少有别的产品添加进这个等级结构
// 大大减少了工厂类的数量

// 抽象层
type AbstractApple interface {
	ShowApple()
}
type AbstractBanana interface {
	ShowBanana()
}
type AbstractPear interface {
	ShowPear()
}
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现层
// 中国产品
type ChineseApple struct{}

func (ca *ChineseApple) ShowApple() {
	fmt.Println("Chinese apple")
}

type ChineseBanana struct{}

func (ca *ChineseBanana) ShowBanana() {
	fmt.Println("Chinese banana")
}

type ChinesePear struct{}

func (ca *ChinesePear) ShowPear() {
	fmt.Println("Chinese pear")
}

type ChinaFac struct{}

func (cf *ChinaFac) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChineseApple)
	return apple
}
func (cf *ChinaFac) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChineseBanana)
	return banana
}
func (cf *ChinaFac) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinesePear)
	return pear
}

// 日本产品
type JapanApple struct{}

func (ca *JapanApple) ShowApple() {
	fmt.Println("Japan apple")
}

type JapanBanana struct{}

func (ca *JapanBanana) ShowBanana() {
	fmt.Println("Japan banana")
}

type JapanPear struct{}

func (ca *JapanPear) ShowPear() {
	fmt.Println("Japan pear")
}

type JapanFac struct{}

func (cf *JapanFac) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(JapanApple)
	return apple
}
func (cf *JapanFac) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(JapanBanana)
	return banana
}
func (cf *JapanFac) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(JapanPear)
	return pear
}

// 业务逻辑层
func AbstractFacPattern() {
	// 生产中国苹果、香蕉、梨
	// 首先生成一个中国工厂
	var cFac AbstractFactory
	cFac = new(ChinaFac)
	// 生成苹果、香蕉、梨
	var cApple AbstractApple
	cApple = cFac.CreateApple()
	cApple.ShowApple()
	var cBanana AbstractBanana
	cBanana = cFac.CreateBanana()
	cBanana.ShowBanana()
	var cPear AbstractPear
	cPear = cFac.CreatePear()
	cPear.ShowPear()

	// 生产日本水果同理
	var jFac AbstractFactory
	jFac = new(JapanFac)
	jBanana := jFac.CreateBanana()
	jBanana.ShowBanana()
}

// 针对”产品族“进行添加，符合开闭原则
// 针对”产品等级结构“进行添加，不符合开闭原则

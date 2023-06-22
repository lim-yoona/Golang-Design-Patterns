package creational

import "fmt"

// 与简单工厂模式不同的是
// 工厂方法模式将简单工厂模式中的一个工厂做一个抽象
// 并依赖这个抽象，有很多具体的工厂，返回相应的对象

// 抽象层
// 水果抽象
type AbstractFruit interface {
	ShowFruit()
}

// 工厂抽象
type AbstractFac interface {
	CreateFruit() AbstractFruit
}

// 基础模块层
// 具体的水果
type TheApple struct{}

func (a *TheApple) ShowFruit() {
	fmt.Println("This is an Apple")
}

type TheBanana struct{}

func (a *TheBanana) ShowFruit() {
	fmt.Println("This is a Banana")
}

type ThePear struct{}

func (a *ThePear) ShowFruit() {
	fmt.Println("This is a Pear")
}

// 具体的工厂
type AppleFac struct{}

func (af *AppleFac) CreateFruit() AbstractFruit {
	var apple AbstractFruit
	apple = new(TheApple)
	return apple
}

type BananaFac struct{}

func (af *BananaFac) CreateFruit() AbstractFruit {
	var banana AbstractFruit
	banana = new(TheBanana)
	return banana
}

type PearFac struct{}

func (af *PearFac) CreateFruit() AbstractFruit {
	var pear AbstractFruit
	pear = new(ThePear)
	return pear
}

// 业务逻辑层
// 业务逻辑层只需要面向抽象层编写代码即可
func FacMethod() {
	// 想产生一个苹果对象
	// 首先创建一个苹果工厂
	var appFac AbstractFac
	appFac = new(AppleFac)

	// 生成苹果
	var apple AbstractFruit
	apple = appFac.CreateFruit()
	apple.ShowFruit()

	// 想产生一个香蕉对象
	var bananaFac AbstractFac
	bananaFac = new(BananaFac)
	var banana AbstractFruit
	banana = bananaFac.CreateFruit()
	banana.ShowFruit()

	// 想要增加一个水果也很简单
	// 增加水果类，增加对应的工厂类，业务逻辑层面向接口编程，符合开闭原则思想
	// 缺点是增加了系统中类的个数，增加了系统的抽象性和理解难度
}

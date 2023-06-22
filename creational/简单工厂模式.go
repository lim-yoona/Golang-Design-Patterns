package creational

import "fmt"

// 为什么需要工厂模式？
// 如果没有工厂，我们的业务逻辑需要依赖基础类模块，耦合度变高
// 并且随着我们期望的创建出来的对象种类的增加，可能需要时常修改构造函数，违反了开闭原则
// 简单工厂模式，构造函数仅依赖抽象类，而不依赖具体的类
// 引入工厂模式层：业务逻辑层--> 工厂模块 -->基础类模块

// ----------抽象层----------
// 抽象出来一个水果接口
type Fruit interface {
	Show()
}

// ----------实现层----------
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("Here Apple")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("Here Banana")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("Here Pear")
}

// 工厂层
type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

// ----------业务逻辑层----------
func Business() {
	fac := new(Factory)

	apple := fac.CreateFruit("apple")
	apple.Show()

	ban := fac.CreateFruit("banana")
	ban.Show()
}

// 缺点：简单工厂模式违反了开闭原则，新增加一个水果的时候需要修改工厂模块
// 缺点：工厂类职责过重，一旦不能工作，系统收到影响

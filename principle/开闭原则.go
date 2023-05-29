package principle

import "fmt"

// 开闭原则是指类的改动是通过增加代码进行的，而不是修改代码
// 开闭原则，可以通过继承来实现，每个功能一个类
// 这样增加新的功能只需要增加新的类，而不需要修改现有的类
type AbstractBank interface {
	Deal()
}

// 功能类实现接口
type SaveBank struct{}

func (sb *SaveBank) Deal() {
	fmt.Println("存款业务...")
}

type TranBank struct{}

func (tb *TranBank) Deal() {
	fmt.Println("转账业务...")
}

// 增加功能
// 只需要增加一个类实现接口，而不用修改之前的类
type GetMoneyBank struct{}

func (gb *GetMoneyBank) Deal() {
	fmt.Println("取钱...")
}

// 这样一来，多态的实现也就是很简单的事情了
func BankBusiness(bank AbstractBank) {
	// 实现多态
	bank.Deal()
}

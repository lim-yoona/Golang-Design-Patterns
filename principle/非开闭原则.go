package principle

import "fmt"

// 非开闭原则的情况下
// 我们想想一个银行类，它应该提供很多的功能
// 比如存款、转账等

type bank struct{}

func (b *bank) save() {
	fmt.Println("存款业务...")
}
func (b *bank) tran() {
	fmt.Println("存款业务...")
}

// 如果这个银行需要加功能，那么就需要修改这个类
// 如果实现这个函数需要增加类的属性
// 那就要修改类属性，可能会影响别的功能
func (b *bank) getMoney() {
	fmt.Println("存款业务...")
}

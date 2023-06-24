package creational

import "fmt"

// 单例模式,保证一个类
// 只有一个实例存在
// 同时提供能对该实例加以访问的全局访问方法

// 首先保证这个类非公有化，外界不能通过这个类直接创建一个对象
// 那么这个类就应该变得非公有访问，类名称首字母小写

// 别的语言中，要通过类的构造函数设置为私有，而别人不能创建对象
// Go由于没有构造函数，可以通过把类名小写，不允许其他包访问这个类
// 并且把单例模式设置为一个单独的包，来实现单例模式
type singleClass struct{}

// 需要有一个指针可以指向这个唯一对象，但是这个指针不能发生改变，如果这个指针设置为大写开头
// Golang中没有常指针，只能通过将指针私有化不让其他包访问
var instance *singleClass = new(singleClass)

// 如果全部都是私有化，那么外部模块将永远无法访问这个对象
// 所需要对外提供一个方法来获取到这个对象
func GetInstance() *singleClass {
	return instance
}
func (sc *singleClass) Output() {
	fmt.Println("我是单例instance")
}

// 这样的单例写法是"饿汉式"的单例模式,就是说，对象一定是会被实例化的

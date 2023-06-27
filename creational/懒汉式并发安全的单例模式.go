package creational

// 与饿汉式单例模式不同，懒汉式单例模式只有在获取对象函数被调用过才会实例化一个对象

import (
	"fmt"
	"sync"
)

type singleClass1 struct{}

var once sync.Once

// 需要有一个指针可以指向这个唯一对象，但是这个指针不能发生改变，如果这个指针设置为大写开头
// Golang中没有常指针，只能通过将指针私有化不让其他包访问
var instance1 *singleClass1

// 如果全部都是私有化，那么外部模块将永远无法访问这个对象
// 所需要对外提供一个方法来获取到这个对象
func GetInstance1() *singleClass1 {
	// 为了确保并发安全，这里采用sync.Once确保Do中的函数只执行一次
	once.Do(func() {
		instance1 = new(singleClass1)
	})
	return instance1
}
func (sc *singleClass1) Output1() {
	fmt.Println("我是懒汉并发安全单例instance")
}

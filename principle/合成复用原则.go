package principle

import "fmt"

// 如果使用继承，会导致所有父类的改变都可能会影响到子类的行为
// 如果使用对象组合，就降低了这种依赖关系
// 因此对于继承和组合都能完成的功能，优先使用组合

// 以一个猫的例子来说明如何运用好合成复用原则

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("Cat Eat...")
}

// 现在这个Cat类有一个Eat方法
// 如果我想给猫添加一些别的方法比如Sleep
// 在遵循开闭原则的情况下，不去修改原Cat类，而是新写一个猫类依赖Cat
// 有两种实现方法，其一是继承，其二是组合

// 继承实现
type CatB struct {
	Cat
}

// 新增方法Sleep
func (cb *CatB) Sleep() {
	fmt.Println("Cat Sleep...")
}

// 组合实现1
type CatC struct {
	c *Cat
}

func (cc *CatC) Eat() {
	cc.c.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("Cat Sleep...")
}

// 组合实现2
// 只让Eat方法依赖Cat类，而不把Cat作为属性
// 依赖更小
type CatD struct{}

func (cc *CatD) Eat(c *Cat) {
	c.Eat()
}

func (cc *CatD) Sleep() {
	fmt.Println("Cat Sleep...")
}

func TestCat() {
	cat := &Cat{}
	// 继承实现
	fmt.Println("继承实现...")
	c := &CatB{}
	c.Eat()
	c.Sleep()

	// 组合实现1
	fmt.Println("组合实现1...")
	cc := &CatC{}
	cc.Eat()
	cc.Sleep()

	// 组合实现2
	fmt.Println("组合实现2...")
	cd := &CatD{}
	cd.Eat(cat)
	cd.Sleep()
}

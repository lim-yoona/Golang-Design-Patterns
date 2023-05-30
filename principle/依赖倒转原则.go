package principle

import "fmt"

// 在设计系统时，需要避免耦合度极高的设计
// 应当将系统分为三层，分别是抽象层、实现层和业务逻辑层
// 抽象层提供抽象的接口，增加实体（类）（即实现层）只需要去实现接口即可，不需要去关注别的类
// 业务逻辑层写业务，只需要关注抽象层中的接口即可，而不需要关注很多类
// 何为依赖倒转原则：业务逻辑层依赖抽象层，实现层依赖抽象层

// 以使用电脑为例，来设计一个系统

// **************抽象层****************
// 设计一个抽象的电脑的接口
type abstractComputer interface {
	// 电脑启动
	StartComputer()
}

// 设计一个抽象的人的接口
type abstractPerson interface {
	// 使用电脑,传入抽象的电脑进去，方便实现多态
	Use(computer abstractComputer)
}

// 抽象层只依赖抽象层，而不去管别的层

// **************实现层****************
// 实现层依赖抽象层
// 一个人名为Jack,实现abstractPerson接口
type Jack struct {
	// 一些属性
}

func (j *Jack) Use(com abstractComputer) {
	// 使用电脑
	fmt.Println("Jack使用电脑...")
	com.StartComputer()
}

// 一个人名为Mary,实现abstractPerson接口
type Mary struct {
	// 一些属性
}

func (m *Mary) Use(com abstractComputer) {
	// 使用电脑
	fmt.Println("Mary使用电脑...")
	com.StartComputer()
}

// 联想牌电脑，实现abstractComputer接口
type lianxiang struct {
	// 一些属性
}

func (lx *lianxiang) StartComputer() {
	fmt.Println("联想电脑启动...")
}

// 惠普牌电脑，实现abstractComputer接口
type hp struct {
	// 一些属性
}

func (lx *hp) StartComputer() {
	fmt.Println("惠普电脑启动...")
}

// 如果说需要增加一个人或者一台电脑，是非常方便的
type yoona struct {
	// 一些属性
}

func (y *yoona) Use(com abstractComputer) {
	// 使用电脑
	fmt.Println("yoona使用电脑...")
	com.StartComputer()
}

// **************业务逻辑层****************
// 业务逻辑层只依赖抽象层
func Business() {
	var com abstractComputer
	var per abstractPerson
	// 要写一个Jack使用联想电脑的代码
	// 定义jack变量
	per = new(Jack)
	com = new(lianxiang)
	per.Use(com)

	// Mary使用联想电脑
	// 定义Mary变量
	per = new(Mary)
	per.Use(com)
	// Mary使用惠普电脑
	com = new(hp)
	per.Use(com)

	// 增加了yoona之后写业务代码
	var p abstractPerson
	p = new(yoona)
	p.Use(com)
}

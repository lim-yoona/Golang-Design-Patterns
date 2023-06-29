package behavioral

import "fmt"

// 装饰器模式，装饰某一个对象，为一个对象增加额外的属性
// 例如，给手机贴膜以及装手机壳

// 抽象层
type Phone interface {
	ShowPhone()
}

// 实现层
// 首先实现两部手机
type Huawei struct{}

func (h *Huawei) ShowPhone() {
	fmt.Println("一部华为手机")
}

type Xiaomi struct{}

func (x *Xiaomi) ShowPhone() {
	fmt.Println("一部小米手机")
}

// 实现装饰器
type Decorator struct {
	phone Phone
}

func (d *Decorator) ShowPhone() {}

// 贴膜装饰器
type MoDecorator struct {
	Decorator
}

func (d *MoDecorator) ShowPhone() {
	d.phone.ShowPhone()
	// 贴膜逻辑
	fmt.Println("贴膜的手机")
}
func NewMoDecorator(p Phone) Phone {
	return &MoDecorator{
		Decorator{
			phone: p,
		},
	}
}

// 手机壳装饰器
type KeDecorator struct {
	Decorator
}

func (k *KeDecorator) ShowPhone() {
	k.phone.ShowPhone()
	// 加壳逻辑
	fmt.Println("带壳的手机")
}
func NewKeDecorator(p Phone) Phone {
	return &KeDecorator{
		Decorator{
			phone: p,
		},
	}
}

// 业务逻辑层
func DecoratorFunc() {
	var huawei Phone
	huawei = new(Huawei)
	huawei.ShowPhone()

	// 给华为手机贴膜
	var MoHuawei Phone
	MoHuawei = NewMoDecorator(huawei)
	MoHuawei.ShowPhone()

	// 给华为手机装壳
	var KeMoHuawei Phone
	KeMoHuawei = NewKeDecorator(MoHuawei)
	KeMoHuawei.ShowPhone()

	xiaomi := new(Xiaomi)
	Kexiaomi := NewKeDecorator(xiaomi)
	Kexiaomi.ShowPhone()
}

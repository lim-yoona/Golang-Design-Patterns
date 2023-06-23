package creational

import "fmt"

//为抽象工厂模式添加一个案例，要求：
//设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，显卡具有显示功能（display，功能实现只要打印出意义即可），内存具有存储功能（storage），cpu具有计算功能（calculate）。
//现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
//要求组装两台电脑，
//
//一台（Intel的CPU，Intel的显卡，Intel的内存）
//一台（Intel的CPU，nvidia的显卡，Kingston的内存）
//用抽象工厂模式实现。

// 显卡、内存、CPU这些产品是不变的，变的只是厂商
// 因此工厂模块应该包括Intel工厂，nvidia工厂以及Kingston工厂
// 电脑采用简单工厂模式

// 抽象层
// 抽象的电脑
type AbstractComputer interface {
	ShowComputer()
}

// 抽象的电脑工厂
type AbstractComputerFac interface {
	CreateComputer(cpu AbstractCPU, vcard AbstractVCard, mem AbstractMem) AbstractComputer
}

// 抽象的显卡
type AbstractVCard interface {
	ShowVCard()
}

// 抽象的内存
type AbstractMem interface {
	ShowMemory()
}

// 抽象的CPU
type AbstractCPU interface {
	ShowCPU()
}

// 抽象的工厂
type AbstractF interface {
	// 提供对于三种产品的构造器
	CreateVCard() AbstractVCard
	CreateMem() AbstractMem
	CreateCPU() AbstractCPU
}

// 实现层
type IntelCPU struct{}

func (ic *IntelCPU) ShowCPU() {
	fmt.Println("This is IntelCPU")
}

type IntelVCard struct{}

func (ivc *IntelVCard) ShowVCard() {
	fmt.Println("This is IntelVCard")
}

type NvidiaVCard struct{}

func (nvc *NvidiaVCard) ShowVCard() {
	fmt.Println("This is NvidiaVCard")
}

type NvidiaMem struct{}

func (nvm *NvidiaMem) ShowMemory() {
	fmt.Println("This is NvidiaMem")
}

type NvidiaCPU struct{}

func (nvc *NvidiaCPU) ShowCPU() {
	fmt.Println("This is NvidiaCPU")
}

type IntelMem struct{}

func (im *IntelMem) ShowMemory() {
	fmt.Println("This is IntelMem")
}

type KingstonMem struct{}

func (km *KingstonMem) ShowMemory() {
	fmt.Println("This is KingstonMem")
}

type KingstonCPU struct{}

func (kc *KingstonCPU) ShowCPU() {
	fmt.Println("This is KingstonCPU")
}

type KingstonVCard struct{}

func (kv *KingstonVCard) ShowVCard() {
	fmt.Println("This is KingstonVCard")
}

// 工厂实现
type IntelFac struct{}

func (iFac *IntelFac) CreateVCard() AbstractVCard {
	var intelVCard AbstractVCard
	intelVCard = new(IntelVCard)
	return intelVCard
}
func (iFac *IntelFac) CreateMem() AbstractMem {
	var intelMem AbstractMem
	intelMem = new(IntelMem)
	return intelMem
}
func (iFac *IntelFac) CreateCPU() AbstractCPU {
	var intelCPU AbstractCPU
	intelCPU = new(IntelCPU)
	return intelCPU
}

type NvidiaFac struct{}

func (nFac *NvidiaFac) CreateVCard() AbstractVCard {
	var nvidiaVCard AbstractVCard
	nvidiaVCard = new(NvidiaVCard)
	return nvidiaVCard
}
func (nFac *NvidiaFac) CreateMem() AbstractMem {
	var nvidiaMem AbstractMem
	nvidiaMem = new(NvidiaMem)
	return nvidiaMem
}
func (nFac *NvidiaFac) CreateCPU() AbstractCPU {
	var nvidiaCPU AbstractCPU
	nvidiaCPU = new(NvidiaCPU)
	return nvidiaCPU
}

type KingstonFac struct{}

func (kFac *KingstonFac) CreateVCard() AbstractVCard {
	var kingVCard AbstractVCard
	kingVCard = new(KingstonVCard)
	return kingVCard
}
func (kFac *KingstonFac) CreateMem() AbstractMem {
	var kingMem AbstractMem
	kingMem = new(KingstonMem)
	return kingMem
}
func (kFac *KingstonFac) CreateCPU() AbstractCPU {
	var kingCPU AbstractCPU
	kingCPU = new(KingstonCPU)
	return kingCPU
}

// 具体的电脑产品
type Computer struct {
	Memory AbstractMem
	VCard  AbstractVCard
	CPU    AbstractCPU
}

func (c *Computer) ShowComputer() {
	fmt.Println("-------------This is computer----------------")
	c.VCard.ShowVCard()
	c.CPU.ShowCPU()
	c.Memory.ShowMemory()
}

// 具体的电脑工厂
type ComputerF struct{}

func (cf *ComputerF) CreateComputer(cpu AbstractCPU, vcard AbstractVCard, mem AbstractMem) AbstractComputer {
	var computer AbstractComputer
	computer = &Computer{
		Memory: mem,
		CPU:    cpu,
		VCard:  vcard,
	}
	return computer
}

// 业务逻辑层
func CumputerTest() {
	// 一台（Intel的CPU，Intel的显卡，Intel的内存）
	var computer1 AbstractComputer
	// 首先创建工厂
	var intelF, nvidiaF, KingstonF AbstractF
	intelF = new(IntelFac)
	nvidiaF = new(NvidiaFac)
	KingstonF = new(KingstonFac)
	var computerFac AbstractComputerFac
	computerFac = new(ComputerF)
	// 创建零件
	intelC := intelF.CreateCPU()
	intelVc := intelF.CreateVCard()
	intelMem := intelF.CreateMem()
	// 组装
	computer1 = computerFac.CreateComputer(intelC, intelVc, intelMem)
	computer1.ShowComputer()
	// 一台（Intel的CPU，nvidia的显卡，Kingston的内存）
	// 创建电脑
	var computer2 AbstractComputer
	// 创建零件
	intelcpu := intelF.CreateCPU()
	nvidiaVc := nvidiaF.CreateVCard()
	kingMem := KingstonF.CreateMem()
	// 组装
	computer2 = computerFac.CreateComputer(intelcpu, nvidiaVc, kingMem)
	computer2.ShowComputer()
}

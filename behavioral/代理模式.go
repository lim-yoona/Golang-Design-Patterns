package behavioral

import "fmt"

// 代理模式，使用一个类来代理完成若干种类似的活动
// 下面将使用海外代购来说明代理模式及其使用

// 抽象层
type shopping interface {
	Buy(p *produce)
}

// 实现层
type produce struct {
	// 产品类别
	kind string
	// 产品真伪
	fact bool
}

// 我们回去多个不同的国家购物
type japanShopping struct{}

func (js *japanShopping) Buy(p *produce) {
	fmt.Println("日本购买" + p.kind)
}

type usShopping struct{}

func (js *usShopping) Buy(p *produce) {
	fmt.Println("美国购买" + p.kind)
}

// 我们去买很不方便，所以我们找代理来完成这些事情
// 代理除了可以帮我们买东西，还可以完成一些额外的事情
type proxy struct {
	// 代理是针对某个特定的任务来进行操作
	shop shopping
}

// 代理除了可以购买东西，还可以过海关以及辨别真伪
func (p *proxy) pass(pro *produce) {
	fmt.Println(pro.kind + "过海关")
}
func (p *proxy) judge(pro *produce) bool {
	fmt.Println("判断真货假货")
	if !pro.fact {
		fmt.Println(pro.kind + "假货,不应该购买")
	}
	return pro.fact
}
func (p *proxy) Buy(pro *produce) {
	// 首先判断是真是假
	if !p.judge(pro) {
		fmt.Println(pro.kind + "假货，没买")
		return
	}
	// 代买东西
	p.shop.Buy(pro)
	// 过海关
	p.pass(pro)
}

func createProxyShop(shop shopping) *proxy {
	return &proxy{
		shop: shop,
	}
}

// 业务逻辑层
func ProxyBuy() {
	// 两件商品
	pro1 := &produce{
		kind: "苹果",
		fact: true,
	}
	pro2 := &produce{
		kind: "宝马",
		fact: false,
	}

	var jShoping shopping
	jShoping = new(japanShopping)
	// 请代购
	proxyS := createProxyShop(jShoping)
	proxyS.Buy(pro1)

	var aShoping shopping
	aShoping = new(usShopping)
	// 请代购
	proxyA := createProxyShop(aShoping)
	proxyA.Buy(pro2)
}

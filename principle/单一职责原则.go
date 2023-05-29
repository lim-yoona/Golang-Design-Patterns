package principle

import "fmt"

// 单一职责原则旨在确保一个类的职责是单一的
// 如果一个类完成多个职责，那么修改其中一个职责，可能会影响对于其他职责语义/逻辑的理解
// 但是如果写成多个类，就会很清晰

// 正确示范
// 一个java工程师
type JavaWorker struct{}

func (jw *JavaWorker) work() {
	fmt.Println("编写Java代码...")
}

// 一个Golang工程师
type GolangWorker struct{}

func (gw *GolangWorker) work() {
	fmt.Println("编写Golang代码...")
}

// 错误示范
type Worker struct{}

func (w *Worker) javaWork() {
	fmt.Println("编写Java代码...")
	// 但java工程师也会被要求写golang代码
	fmt.Println("编写Golang代码...")
}
func (w *Worker) golangWork() {
	fmt.Println("编写Golang代码...")
	// 但java工程师也会被要求写java代码
	fmt.Println("编写Java代码...")
}

// 由于类Worker的职责并不单一
// 如果javaWork函数和golangWork函数实现类似
// 可能会使代码可读性变差，且逻辑混乱，易混淆

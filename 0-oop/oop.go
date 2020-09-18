package __oop

import "fmt"

// Golang使用结构体模拟类与对象
type Bike struct {
	color string
	Name  string
}

func (b *Bike) GetColor() string {
	return b.color
}

// 封装：首字母大小写代表共有私有权限
type Person struct {
	name string
}

func (p *Person) Walk() string {
	return p.name
}

// 继承：使用内嵌的方式，对结构体struct进行组合，利用匿名结构体
type Chinese struct {
	p    Person
	skin string
}

func (c *Chinese) GetSkin() string {
	return "yellow"
}

// 多态 : 使用interface来实现
type Human interface {
	Speak()
}

type American struct {
	name     string
	Language string
}

func (a *American) Speak() {
	fmt.Println(a.name + "balabala")
}

type Ateman struct {
	name   string
	Height int
}

func (a *Ateman) Speak() {
	fmt.Println(a.name + "balabala")
}

//========================================================================
//   违反五大原则
//
//========================================================================

//------------------------------------------------------------------------
// 违反单一功能原则：By方法里面又有Move方法
// 违反开闭原则：此时要增加摩托车的时候，就要再增加一个case
func (p *Person) By(name string) {
	switch name {
	case "bike":
		bike := Bike{}
		bike.Move()
	case "car":
		car := Car{}
		car.Move()
	}
}

// 应该这么改：要增加摩托车的时候，就在增加一个function就好了
func (p *Person) RideBike(bike Bike) {
	bike.Move()
}

func (p *Person) DriveCar(car Car) {
	car.Move()
}

//------------------------------------------------------------------------

//------------------------------------------------------------------------
type V6Engine struct {
}

func (v6 *V6Engine) Run() string {
	return "V6Engine run"
}

// 这种方式不好，它依赖于具体的实现
type BMWCar struct {
	engine V6Engine
}

func (car *BMWCar) RunEngine() string {
	return "BMWCar的" + car.engine.Run()
}

// 应该这么改：
type engine interface {
	Run() string
}

type BMWEngine struct {
	engine engine
}

func (car *BMWCar) RunEngine() string {
	return "BMWCar的" + car.engine.Run()
}

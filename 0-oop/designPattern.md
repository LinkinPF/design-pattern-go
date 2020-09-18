### 为什么学习设计模式？

解决代码的解耦、可复用性、扩展性、接口隔离、单一功能原则等等

### golang是不是面向对象的语言？

不是，但是golang可以面向对象编程。

所以首先解决掉goalng的oop，然后上设计模式

## golang 的 oop
主要分为下面三个部分
- 类与对象
- 三大基本特性
- 五大基本原则

### golang 的类与对象

Golang使用结构体模拟类与对象
```go
type Bike struct {
	color string
	Name  string
}

func (b *Bike) GetColor() string {
	return b.color
}
```

### golang 的三大基本特性
- 封装：首字母大小写代表共有私有权限
- 继承：使用内嵌的方式，对结构体struct进行组合，利用匿名结构体
- 多态：也就是接口的多种实现方式，使用interface来实现

```go
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


//多态：使用interface来实现
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
```


## 五大基本原则
- 单一功能原则：也就是一个对象的方法的功能要单一，要走路就走路，要跳就跳，不要又能走又能跳的
- 开闭原则：对扩展是开放的，对修改是关闭的
- 里氏替换原则：子类可以替换父类，对于go语言来说，就是要面向接口编程
- 接口隔离原则：接口最小化
- 依赖反转原则：不依赖具体实现，要面向接口开发

有几个例子可以看下oop.go


## 什么是设计模式
设计模式就是一套被反复使用、多人知晓的、经过分类的、代码设计经验的总结







package __simple_factory

import "fmt"

/*
	原理：https://design-patterns.readthedocs.io/zh_CN/latest/creational_patterns/simple_factory.html

	这个包中只有 API 和 NewAPI 是包外可见的，封装了实现的细节
*/

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

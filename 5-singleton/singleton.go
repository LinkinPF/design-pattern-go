package __singleton

import "sync"

//这个是单例模式类
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 方法用来获取单例模式类对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

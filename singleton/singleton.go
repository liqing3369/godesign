package singleton

import "sync"

// Singleton 创建目标
type Singleton struct{}

var singleton *Singleton
var once sync.Once

// GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

package lazy

import (
	"go-study/design_pattern/singleton"
	"sync"
)

var (
	instance *singleton.Singleton
	once     sync.Once
)

// GetInstance
//  @Description: 懒汉式单例模式
//  @return *Singleton
func GetInstance() *singleton.Singleton {
	once.Do(func() {
		instance = new(singleton.Singleton)
	})

	return instance
}

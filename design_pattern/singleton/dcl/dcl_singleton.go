package dcl

import (
	"go-study/design_pattern/singleton"
	"sync"
)

var (
	instance *singleton.Singleton
	mu       sync.Mutex
)

// GetInstance
//  @Description: 饿汉式单例模式
//  @return *singleton.Singleton
func GetInstance() *singleton.Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(singleton.Singleton)
		}
	}

	return instance
}

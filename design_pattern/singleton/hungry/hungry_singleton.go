package hungry

import "go-study/design_pattern/singleton"

var instance = new(singleton.Singleton)

// GetInstance
//  @Description: 饿汉式单例模式
//  @return *singleton.Singleton
func GetInstance() *singleton.Singleton {
	return instance
}

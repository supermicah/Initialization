package design

import (
	"fmt"
	"sync"
)

var once sync.Once

type Singleton struct{}

var instance *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
	})
	return instance
}

func (s *Singleton) SomeThing() {
	fmt.Println("单例对象的某方法")
}
func do3d4() {
	s := GetInstance()
	s.SomeThing()
}

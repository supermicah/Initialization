package main

import "fmt"

// Fruit 水果类(抽象接口)
type Fruit interface {
	Show() //接口的某方法
}

// AbstractFactory 工厂类(抽象接口)
type AbstractFactory interface {
	CreateFruit() Fruit //生产水果类(抽象)的生产器方法
}

type Apple struct {
	Fruit
}

func (f *Apple) Show() {
	fmt.Println("我是苹果")
}

type AppleFactory struct {
	AbstractFactory
}

func (f *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

func do3d1() {
	factory := new(AppleFactory)
	fruit := factory.CreateFruit()
	fruit.Show()

}

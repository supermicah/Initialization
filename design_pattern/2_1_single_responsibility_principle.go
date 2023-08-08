package main

//单一职责原则

import "fmt"

type ClothesShop struct{}

func (cs *ClothesShop) OnShop() {
	fmt.Println("休闲的装扮")
}

type ClothesWork struct{}

func (cw *ClothesWork) OnWork() {
	fmt.Println("工作的装扮")
}

func do2d1() {
	//工作的时候
	cw := new(ClothesWork)
	cw.OnWork()

	//shopping的时候
	cs := new(ClothesShop)
	cs.OnShop()
}

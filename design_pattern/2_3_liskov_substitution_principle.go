package main

import "fmt"

type Car interface {
	Run()
}

// Driver 驾驶员信息
//type Driver interface {
//	Drive(c Car)
//}

type Benz struct {
}

func (c *Benz) Run() {
	fmt.Println("Benz is Running.")
}

type BWM struct {
}

func (c *BWM) Run() {
	fmt.Println("BWM is Running.")
}

type ZhangSan struct {
}

func (d *ZhangSan) Drive(c Car) {
	fmt.Println("ZhangSan is Driving.")
	c.Run()
}

type LiSi struct {
}

func (d *LiSi) Drive(c Car) {
	fmt.Println("LiSi is Driving.")
	c.Run()
}

func do2d3() {
	bwm := BWM{}
	zhang3 := &ZhangSan{}
	zhang3.Drive(&bwm)

	benz := Benz{}
	lisi := &LiSi{}
	lisi.Drive(&benz)
}

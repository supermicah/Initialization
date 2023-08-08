package main

// 合成复用原则

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

type BCat struct {
	Cat
}

func (c *BCat) Sleep() {
	fmt.Println("小猫睡觉")
}

type CCat struct {
	C *Cat
}

func (c *CCat) Sleep() {
	fmt.Println("小猫睡觉")
}

func do2d4() {
	//通过继承增加的功能，可以正常使用
	b := BCat{}
	b.Eat()
	b.Sleep()

	//通过组合增加的功能，可以正常使用
	c := CCat{}
	c.C = &Cat{}
	c.C.Eat()
	c.Sleep()
}

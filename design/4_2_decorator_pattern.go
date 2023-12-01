package design

import "fmt"

// Phone 抽象的构件
type Phone interface {
	Show() //构件的功能
}

// Decorator 装饰器基础类（该类本应该为interface，但是Golang interface语法不可以有成员属性）
type Decorator struct {
	phone Phone
}
type HuaWei struct{}

func (hw *HuaWei) Show() {
	fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct{}

func (xm *XiaoMi) Show() {
	fmt.Println("秀出了XiaoMi手机")
}

type MoDecorator struct {
	Decorator //继承基础装饰器类(主要继承Phone成员属性)
}

func (md *MoDecorator) Show() {
	md.phone.Show()      //调用被装饰构件的原方法
	fmt.Println("贴膜的手机") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator //继承基础装饰器类(主要继承Phone成员属性)
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("手机壳的手机") //装饰额外的方法
}

// NewKeDecorator decorator
func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

func do4d2() {
	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show()

	moHuawei := NewMoDecorator(huawei)
	moHuawei.Show()
}

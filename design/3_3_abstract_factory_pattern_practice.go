package design

import "fmt"

type AbstractDisplay interface {
	Display()
}

type AbstractStorage interface {
	Storage()
}

type AbstractCalculate interface {
	Calculate()
}

type IntelDisplay struct{}

func (ca *IntelDisplay) Display() {
	fmt.Println("Inter Display")
}

type IntelStorage struct{}

func (cb *IntelStorage) Storage() {
	fmt.Println("Intel Storage")
}

type IntelCalculate struct{}

func (cp *IntelCalculate) Calculate() {
	fmt.Println("Intel Calculate")
}

type IntelFactory struct{}

func (cf *IntelFactory) Display() AbstractDisplay {
	var d AbstractDisplay

	d = new(IntelDisplay)

	return d
}

func (cf *IntelFactory) Storage() AbstractStorage {
	var s AbstractStorage

	s = new(IntelStorage)

	return s
}

func (cf *IntelFactory) Calculate() AbstractCalculate {
	var pear AbstractCalculate

	pear = new(IntelCalculate)

	return pear
}

func do3d3p() {
	factory := new(IntelFactory)
	factory.Display().Display()

	factory.Storage().Storage()
	factory.Calculate().Calculate()
}

package receiver

import "fmt"

type PersonValue struct {
	name string
}

type PersonPoint struct {
	name string
}

func (p PersonValue) change(name string) {
	p.name = name
}

func (p *PersonPoint) change(name string) {
	p.name = name
}

func Run() {
	pv := PersonValue{name: "personValue"}
	fmt.Printf("%+v\n", pv)
	pv.change("changePersonValue")
	fmt.Printf("%+v\n", pv)

	pp := &PersonPoint{name: "personPoint"}
	fmt.Printf("%+v\n", *pp)
	pp.change("changePersonPoint")
	fmt.Printf("%+v\n", *pp)

}

package main

type S struct {
	Name string
}

func (s *S) ff() {

}

func f(x interface{}) {

}

func g(x *interface{}) {
}

func main() {
	s := S{Name: ""}
	p := &s
	f(s)
	f(p)
}

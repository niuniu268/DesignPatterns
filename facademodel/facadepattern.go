package main

import "fmt"

type SystemA struct {
}

func (a SystemA) MethodA() {
	fmt.Println("Method A")
}

type SystemB struct {
}

func (b SystemB) MethodB() {
	fmt.Println("Method B")
}

type SystemC struct {
}

func (c SystemC) MethodC() {
	fmt.Println("Method C")
}

// facade pattern offers a facade class in which an interface is created
type Facade struct {
	a *SystemA
	b *SystemB
	c *SystemC
}

func (f Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f Facade) MethodTwo() {
	f.b.MethodB()
	f.c.MethodC()
}

func main() {

	facade := &Facade{
		a: new(SystemA),
		b: new(SystemB),
		c: new(SystemC),
	}

	facade.MethodOne()

	fmt.Println("========")

	facade.MethodTwo()

}

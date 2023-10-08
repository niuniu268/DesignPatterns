package main

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("cat eating")
}

type Pussy1 struct {
	Cat
}

func (p *Pussy1) Sleep() {
	fmt.Println("Pussy sleeping")
}

type Pussy2 struct {
}

func (p2 Pussy2) Eat(c *Cat) {

	c.Eat()

}

func (p2 Pussy2) Sleep() {
	fmt.Println("pussy sleeping")
}

func main() {

	c := &Cat{}
	c.Eat()
	//

	p := &Pussy1{}
	p.Eat()
	p.Sleep()

	p2 := &Pussy2{}
	p2.Eat(&Cat{})
	p2.Sleep()

}

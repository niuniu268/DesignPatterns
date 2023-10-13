package main

import "fmt"

// abstract class includes a temple in order to fulfill the procedure
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddBag()
}

// encapsulate a procedure to fulfill

type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {

	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	t.b.AddBag()

}

type Coffee struct {
	template
}

func (c *Coffee) BoilWater() {
	fmt.Println("make coffee to boil water")

}
func (c *Coffee) Brew() {
	fmt.Println("make coffee to brew")

}
func (c *Coffee) PourInCup() {
	fmt.Println("make coffee to pour in water")

}
func (c *Coffee) AddBag() {
	fmt.Println("make coffee to add bag")

}

func MakeCoffee() *Coffee {

	c := new(Coffee)
	c.b = c

	return c

}

func main() {

	c := MakeCoffee()

	c.MakeBeverage()

}

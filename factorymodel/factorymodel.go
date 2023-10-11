package main

import "fmt"

// abstract level
type Fruit interface {
	Show()
}

// fulfill level
type Apple struct {
	Fruit
}

func (a *Apple) Show() {

	fmt.Println("apple")

}

type Banana struct {
	Fruit
}

func (b *Banana) Show() {

	fmt.Println("banana")

}

type Pear struct {
	Fruit
}

func (p *Pear) Show() {
	fmt.Println("pear")

}

type Factory struct {
}

func (f *Factory) CreateFruit(category string) Fruit {
	var fruit Fruit
	if category == "Apple" {
		fruit = new(Apple)
	} else if category == "Banana" {
		fruit = new(Banana)
	} else if category == "Pear" {
		fruit = new(Pear)
	}
	return fruit
}

//logic level

func main() {

	factory := new(Factory)
	apple := factory.CreateFruit("Apple")
	apple.Show()
	banana := factory.CreateFruit("Banana")
	banana.Show()
	pear := factory.CreateFruit("Pear")
	pear.Show()

}

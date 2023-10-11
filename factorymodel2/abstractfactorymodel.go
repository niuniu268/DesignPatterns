package main

import "fmt"

type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit(name string) *Fruit
}

type Apple struct {
	Fruit
}

func (a *Apple) Show() {

	fmt.Println("Apple")

	fmt.Println(&a)

}

type Banana struct {
	Fruit
}

func (b *Banana) Show() {
	fmt.Println("Banana")

	fmt.Println(&b)

}

type Pear struct {
	Fruit
}

func (p *Pear) Show() {

	fmt.Println("Pear")
	fmt.Println(&p)

}

type AppleFactory struct {
	AbstractFactory
}

func (receiver *AppleFactory) CreateFruit(name string) *Fruit {

	var fruit Fruit

	fruit = new(Apple)

	return &fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (receiver *BananaFactory) CreateFruit(name string) *Fruit {

	var fruit Fruit

	fruit = new(Banana)

	return &fruit
}

type PearFactory struct {
	AbstractFactory
}

func (receiver *PearFactory) CreateFruit(name string) *Fruit {

	var fruit Fruit

	fruit = new(Pear)

	return &fruit
}

func main() {

	var appleF AbstractFactory
	appleF = new(AppleFactory)

	apple := *appleF.CreateFruit("Apple")
	apple.Show()

	var bananaF AbstractFactory
	bananaF = new(BananaFactory)

	banana := *bananaF.CreateFruit("Banana")
	banana.Show()

	var pearF AbstractFactory
	pearF = new(PearFactory)

	pear := *pearF.CreateFruit("Pear")
	pear.Show()

}

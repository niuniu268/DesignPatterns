package main

import "fmt"

// Abstract level
type Vehicle interface {
	Run()
}

type Benz struct {
}

func (b Benz) Run() {
	fmt.Println("benz runs")
}

type BMW struct {
}

func (bm BMW) Run() {
	fmt.Println("BMW runs")
}

type Driver interface {
	Drive(v Vehicle)
}

type A1 struct {
}

func (a A1) Drive(car Vehicle) {

	fmt.Println("A1 drive")

	car.Run()

}

type A2 struct {
}

func (a A2) Drive(car Vehicle) {

	fmt.Println("A2 drive")
	car.Run()

}

//Object level

func main() {
	//Logic level

	a1 := A1{}
	a2 := A2{}

	a1.Drive(Benz{})
	a2.Drive(BMW{})

}

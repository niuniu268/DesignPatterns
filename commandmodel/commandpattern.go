package main

import (
	"fmt"
)

type Doctor struct {
}

func (d Doctor) TreatEyes() {
	fmt.Println("treat eyes")
}

func (d Doctor) TreatNose() {
	fmt.Println("treate nose")
}

type Description interface {
	Treat()
}
type DescriptionEyes struct {
	*Doctor
}

func (e DescriptionEyes) Treat() {
	e.TreatEyes()
}

type DescriptionNose struct {
	*Doctor
}

func (n DescriptionNose) Treat() {
	n.TreatNose()
}

type Nurse struct {
	list []Description
}

func (n Nurse) Notify() {
	for index, cmd := range n.list {
		fmt.Println(index)
		cmd.Treat()

	}
}
func main() {

	d := new(Doctor)
	nose := DescriptionNose{
		Doctor: d,
	}
	nose.Treat()
	eyes := DescriptionEyes{
		Doctor: d,
	}
	eyes.Treat()

	n := new(Nurse)
	n.list = append(n.list, &nose)
	n.list = append(n.list, &eyes)
	n.Notify()

}

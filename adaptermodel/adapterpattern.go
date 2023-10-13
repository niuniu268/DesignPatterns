package main

import "fmt"

type Voltage5 interface {
	Use5v()
}

func Use5v() {
	fmt.Println("apply 5 voltage")

}

type Voltage220 struct {
}

func (v Voltage220) User220v() {
	fmt.Println("applu 220 voltage")

}

type Adapter struct {
	v220 *Voltage220
}

func (a Adapter) Use5v() {
	fmt.Println("apply adapter")
	fmt.Println("use 5v")
	a.v220.User220v()
}

func NewAdapter(v220 *Voltage220) *Adapter {
	return &Adapter{v220: v220}
}

type Phone struct {
	v Voltage5
}

func NewPhone(v Voltage5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("phone is charging")
	p.v.Use5v()

}

func main() {

	phone := NewPhone(NewAdapter(new(Voltage220)))
	phone.Charge()

}

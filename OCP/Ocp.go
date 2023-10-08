package main

import "fmt"

type Banker interface {
	Business()
}

type Banker1 struct {
}

func (b *Banker1) Business() {

	fmt.Println("Payment function")

}

type Banker2 struct {
}

func (b *Banker2) Business() {

	fmt.Println("Transfer function")

}

type Banker3 struct {
}

func (b *Banker3) Business() {

	fmt.Println("Stock function")

}

func BankBusiness(abstractBanker Banker) {
	abstractBanker.Business()

}

func main() {
	//
	//b1 := Banker1{}
	//
	//b1.Business()
	//
	//b2 := Banker2{}
	//
	//b2.Business()
	//
	//b3 := Banker3{}
	//
	//b3.Business()

	BankBusiness(&Banker1{})
	BankBusiness(&Banker2{})
	BankBusiness(&Banker3{})

}

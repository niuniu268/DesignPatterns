package main

import "fmt"

type Goods struct {
	Category    string
	Authentical bool
}
type Purchase interface {
	Buy(goods *Goods)
}

type KoreanPurchase struct {
	Purchase
}

func (k *KoreanPurchase) Buy(goods *Goods) {
	fmt.Println("buy stuffs in Korean: ", goods.Category)

}

type AmericanPurchase struct {
	Purchase
}

func (u *AmericanPurchase) Buy(goods *Goods) {
	fmt.Println("buy stuffs in US: ", goods.Category)

}

type AfricanPurchase struct {
	Purchase
}

func (a *AfricanPurchase) Buy(goods *Goods) {
	fmt.Println("buy stuffs in Africa: ", goods.Category)

}

type OverseasProxy struct {
	shopping Purchase
}

func NewProxy(shopping Purchase) Purchase {
	return &OverseasProxy{shopping}
}

func (op *OverseasProxy) Buy(goods *Goods) {
	//check authenticate
	if op.distinguish(goods) {
		//coordinate the proxy and apply buy method
		op.shopping.Buy(goods)
		//custom check
		op.custom(goods)
	}

}

func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("applying to check ", goods.Category)
	if goods.Authentical != true {
		fmt.Println("there is fake ", goods.Category)

	} else {
		fmt.Println("there is authenitcal", goods.Category)
	}
	return goods.Authentical

}

func (op *OverseasProxy) custom(goods *Goods) {
	fmt.Println("apply custom check ", goods.Category)

}

func main() {

	var g1 = &Goods{
		Category:    "product",
		Authentical: true,
	}

	//g2 := &Goods{
	//	Category:    "product 2",
	//	Authentical: false,
	//}
	var shopping Purchase
	shopping = new(KoreanPurchase)
	shopping.Buy(g1)

	fmt.Println("============")

	var proxyShopping Purchase
	proxyShopping = NewProxy(shopping)
	proxyShopping.Buy(g1)

}

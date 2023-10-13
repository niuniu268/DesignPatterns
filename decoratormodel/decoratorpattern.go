package main

import "fmt"

// abstractive level

type PhoneFunction interface {
	Show()
}

// abstract decorator level
// the decorator should be abstract
// but Golang's interface does not have members' attribute

type Decorator struct {
	function PhoneFunction
}

func (d *Decorator) Show() {

	fmt.Println("decorator function")
}

// project fulfilled level

type Huawei struct {
}

func (h *Huawei) Show() {
	fmt.Println("manifest by huawei cellphone")
}

type Xiaomi struct {
}

func (x *Xiaomi) Show() {
	fmt.Println("manifest by xiaomi cellphone")
}

// specific decorator
//
//type SpecificDecorator struct {
//	//inherit decorator
//	function Decorator
//}
//
//func NewDecorator(function PhoneFunction) PhoneFunction {
//
//	return &SpecificDecorator{Decorator{function: function}}
//}
//
//func (sd *SpecificDecorator) Show() {
//	sd.function.function.Show()
//
//	//	add a new function
//	fmt.Println("add a new function")
//
//}

type XiaomiDecorator struct {
	Decorator
}

func (x *XiaomiDecorator) Show() {

	x.function.Show()

	fmt.Println("add a xiaomi function")

}

func XiaomiDecoratorFunc(function PhoneFunction) PhoneFunction {

	return &XiaomiDecorator{Decorator{function: function}}

}

type HuaweiDecorator struct {
	Decorator
}

func (h *HuaweiDecorator) Show() {

	h.function.Show()

	fmt.Println("add a huawei function")

}

func HuaweiDecoratorFunc(function PhoneFunction) PhoneFunction {

	return &HuaweiDecorator{Decorator{function: function}}

}
func main() {

	huawei := new(Huawei)
	xiaomi := new(Xiaomi)

	huawei.Show()
	xiaomi.Show()

	fmt.Println("=========")

	HuaweiDecoratorFunc(huawei).Show()
	XiaomiDecoratorFunc(xiaomi).Show()

}

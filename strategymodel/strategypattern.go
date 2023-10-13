package main

import "fmt"

//
//type Role struct {
//	equipment string
//}
//
//func (r Role) Battle() {
//	if r.equipment == "Knife" {
//		k := new(Knife)
//		fmt.Println(r.equipment)
//		k.UseWeapon()
//
//	} else {
//		a := new(Ak47)
//		fmt.Println(r.equipment)
//		a.UseWeapon()
//	}
//
//}

type Strategy interface {
	UseWeapon()
}

type Knife struct {
}

func (k *Knife) UseWeapon() {
	fmt.Println("use knife")

}

type Ak47 struct {
}

func (a *Ak47) UseWeapon() {
	fmt.Println("use AK47")
}

type Hero struct {
	strategy Strategy
}

func (h *Hero) SetStrategy(s Strategy) {
	h.strategy = s

}

func (h *Hero) Battle() {
	h.strategy.UseWeapon()
}

func main() {

	//r := &Role{
	//	equipment: "Knife",
	//}
	//r.Battle()

	var h = &Hero{}

	h.SetStrategy(new(Knife))
	h.Battle()

}

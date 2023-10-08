package main

import "fmt"

type ClothShop struct {
}

func (s *ClothShop) Style() {

	fmt.Println("clothes during shopping")

}

type ClothWork struct {
}

func (w *ClothWork) Style() {
	fmt.Println("clothes during working")

}
func main() {

	fmt.Println("during shopping")
	cs := ClothShop{}
	cs.Style()

	fmt.Println("during working")
	ws := ClothWork{}
	ws.Style()

}

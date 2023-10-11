package main

import "fmt"

//guarantee a class only have one instance
//meanwhile, the instance is access by other classes

//class is not public

type singleton struct {
}

func (s *singleton) ShowFunc() {
	fmt.Println("instance's method")
}

//there is a point

var instance *singleton = new(singleton)

// there is a public method by which others can get access to the instance

func GetInstance() *singleton {
	return instance

}
func main() {

	s1 := GetInstance()
	s1.ShowFunc()

	s2 := GetInstance()

	if s1 == s2 {

		fmt.Println("s1==s2")

	}
}

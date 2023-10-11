package main

import (
	"fmt"
	"sync"
)

//guarantee a class only have one instance
//meanwhile, the instance is access by other classes

//class is not public

type singleton struct {
}

func (s *singleton) ShowFunc() {
	fmt.Println("instance's method")
}

// there is a point
//
// var lock sync.Mutex
var instance *singleton

var once sync.Once

//var initialized uint32

// there is a public method by which others can get access to the instance

func GetInstance() *singleton {

	//if atomic.LoadUint32(&initialized) == 1 {
	//	return instance
	//}
	//
	//lock.Lock()
	//defer lock.Unlock()

	once.Do(func() {
		instance = new(singleton)
	})
	//
	//if instance == nil {
	//	instance = new(singleton)
	//	return instance
	//}
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

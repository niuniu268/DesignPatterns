package main

import "fmt"

type Listener interface {
	TeacherComing()
}

type Notify interface {
	AddListener(l Listener)
	RemoveListener(l Listener)
	Notify()
}

type Student1 struct {
	distract string
}

func (s Student1) TeacherComing() {
	fmt.Println("student 1 stop ", s.distract)

}

type Student2 struct {
	distract string
}

func (s Student2) TeacherComing() {
	fmt.Println("student 2 stop ", s.distract)

}

type Master struct {
	list []Listener
}

func (m *Master) AddListener(l Listener) {
	m.list = append(m.list, l)

}

func (m *Master) RemoveListener(l Listener) {
	for index, value := range m.list {

		if value == l {
			m.list = append(m.list[:index], m.list[index+1:]...)
			break
		}

	}
}
func (m *Master) Notify() {
	for _, value := range m.list {

		fmt.Println(value)
		value.TeacherComing()

	}

}

func main() {

	s1 := &Student1{
		distract: "aaa",
	}
	s2 := &Student2{
		distract: "bbb",
	}

	master := new(Master)
	master.AddListener(s1)
	master.AddListener(s2)

	fmt.Println("-------")
	master.Notify()

}

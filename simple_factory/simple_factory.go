package simple_factory

import "fmt"

type Person interface {
	Say()
}

type Man struct {
	Person
}

func (m *Man) Say() {
	fmt.Println("man")
}

type Women struct {
	Person
}

func (w *Women) Say() {
	fmt.Println("women")
}

func NewPerson(t int) Person {
	switch t {
	case 0:
		return &Women{}
	case 1:
		return &Man{}
	}
	return nil
}

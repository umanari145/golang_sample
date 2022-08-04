package sample

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	hoge()
}

func TestDummy(t *testing.T) {
	fmt.Println("dummy")
	hoge()
}

func TestAnimal(t *testing.T) {
	animal := Animal{}

	animal.Name = "test"
	animal.Age = 22
	animal.Pref = "東京"

	animal.toShow()
}

func TestInterface(t *testing.T) {
	sampleInterface()
}

func Testf1(t *testing.T) {
	var a, b = f1()
	fmt.Println(a)
	fmt.Println(b)
}

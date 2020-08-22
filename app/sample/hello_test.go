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

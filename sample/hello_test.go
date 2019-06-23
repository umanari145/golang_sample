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

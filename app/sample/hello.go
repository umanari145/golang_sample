package sample

import (
	"fmt"
	"os"

	"../cusError"
	mainhoge "../modeldir"
	subhoge "../modeldir/subidr"
)

func hoge() {
	fmt.Println("filename main packge main")
	mainhoge.SampleA()
	mainhoge.Sample2()

	subhoge.TestB()
	subhoge.TestB2()

	//参照渡ししなくてもmap自体が参照なので値が変わっている
	arrlist := getMapData()
	fmt.Println(arrlist)
	refCheck(arrlist)
	fmt.Println(arrlist)

	var a1 Animal
	a1.Name = "山田"
	a1.Age = 32
	fmt.Println(a1)
	norefCheck(a1)
	fmt.Println(a1)

	refCheck2(&a1)
	fmt.Println(a1)

	file, err := os.Open("hoge.txt")
	//独自のエラー処理を行う
	if err != nil {
		cusError.Error2()
	}

	fmt.Println(file)
}

func getMapData() []map[string]string {
	m := map[string]string{}
	m["name"] = "松本"
	m["age"] = "37"
	m["pref"] = "千葉"

	m2 := map[string]string{}
	m2["name"] = "佐藤"
	m2["age"] = "39"
	m2["pref"] = "東京"

	//mapのリスト型
	var map_list2 []map[string]string
	map_list2 = append(map_list2, m)
	map_list2 = append(map_list2, m2)

	return map_list2
}

func getClass() []Animal {
	var animalList []Animal
	var a1 Animal
	a1.Name = "山田"
	a1.Age = 32
	animalList = append(animalList, a1)
	animalList = append(animalList, Animal{Name: "鈴木", Age: 22})
	animalList = append(animalList, Animal{Name: "田中", Age: 25})

	return animalList
}

func refCheck(map_list2 []map[string]string) {

	fmt.Println(map_list2)
	//mapはもともと参照を渡している
	for _, mapdata := range map_list2 {
		mapdata["age"] = "33"
	}
	fmt.Println(map_list2)
}

func norefCheck(a Animal) {
	a.Age = 99
}

func refCheck2(a *Animal) {
	a.Age = 99
}

//Animal は動物の構造体
type Animal struct {
	Name string
	Age  int
	Pref string
}

func (animal Animal) toShow() {
	fmt.Printf("name:%s age:%d Pref:%s", animal.Name, animal.Age, animal.Pref)
}

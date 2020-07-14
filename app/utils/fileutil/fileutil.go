package fileutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func loadFile() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	// 出力
	fmt.Println(string(b))
}

func loadJson() {
	f, err := os.Open("json_sample.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	// 出力
	//jsonStr := string(b)

	//fmt.Println(jsonStr)
	var sList samples

	err2 := json.Unmarshal(b, &sList)
	if err2 != nil {
		log.Fatal(err2.Error())
		log.Fatal("error")
	}
	fmt.Println(sList)
}

type samples []sample

type sample struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"ip_address"`
}

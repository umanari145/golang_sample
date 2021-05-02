package httputil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Response is response
type Response struct {
	//エラー発生時にメッセージあり
	Message string `json:"message"`
	//実際のデータがここに入ってくる
	Results []Result `json:"result"`
}

//Result は実際の県データ
type Result struct {
	PrefCode int    `json:"prefCode"`
	PrefName string `json:"prefName"`
}

const (
	//APIURL は都道府県API
	APIURL = "https://opendata.resas-portal.go.jp/api/v1/prefectures"
)

//Doreq Access Http request and parse JSON
func Doreq() {

	request, err := http.NewRequest("GET", APIURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	//クエリパラメータ GETのクエリを入れるときは↓のように
	//params := request.URL.Query()
	//params.Add("userId", "1")
	//request.URL.RawQuery = params.Encode()

	//HTTPヘッダーを入れるときは下記のように
	APIKEY := os.Getenv("APIKEY")
	request.Header.Set("X-API-KEY", APIKEY)

	//fmt.Println(request.URL.String())

	client := http.Client{}

	httpRes, err := client.Do(request)
	if httpRes.StatusCode != 200 && err != nil {
		log.Fatal(err)
	}

	defer httpRes.Body.Close()

	body, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	if response.Message != "" {
		log.Fatal("jsonのパースに失敗しました。")
	}

	for _, result := range response.Results {
		fmt.Printf("都道府県コード:%d 都道府県名:%s", result.PrefCode, result.PrefName)
	}

}

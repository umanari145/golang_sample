# golang_sample

## 特徴

- 並行処理
- 高速コンパイル
- 並列処理に強い


## sandbox
https://play.golang.org/
### 基礎文法

https://www.amazon.co.jp/gp/product/B07HDMR6S5/ref=ku_mi_rw_edp_ku<br>
参照

####  スコープ

```
// グローバル変数
var a int = 999
func main(){
    // ローカル変数
    b := 1
    fmt.Println(a + b)
}

```

### 配列とスライスの違い

同一の型の値を「予め決められた個数分」保持するためのデータ<br>
スライスは「可変長の配列」<br>
配列との違いは
- 定義時にサイズを指定しない
- 要素のサイズが変更できる
- ゼロ値が"nil"※後述


```
// 配列の宣言
var a [3]string    
b := [2]int{999, 888}    
a[0] = "hoge"
```

```
// スライスの宣言
var a []string    
b := []int{999, 888}

```

### 構造体

```
// 構造体の宣言
type MyStruct struct{
    a string    
    b int
}

type MyStruct2 struct{
    MyStruct // 構造体のネストが可能
    c int
}

var st MyStruct
st.a = "hoge"    
fmt.Println(st.a)
var st2 MyStruct2    
st2.a = "fuga" // 参照元の構造体メンバへのアクセス    
st2.c = 999    
fmt.Println(st2.a, st2.c)

```

### ポインタ

```
var p *int// int型の変数
a := 10// 変数aのアドレスを取得
p = &a
fmt.Println(p)
```

### 型変換

```
var a int = 999
var b int16 = 1
var c int = a + int(b)    
fmt.Println(c)
```

### 独自の型
```
# intをベースにした独自の型で変換ができない
type MyInt int

```


### interface
任意の型を取れる

```
var obj interface{}

obj = 1
fmt.Println(obj)

obj = "hoge"
fmt.Println(obj)

obj = []string{"honda", "toyota", "mazda"}
fmt.Println(obj)
```

### errのよく出る書き方

```
err := doError()
if err != nil {
  // エラー時の処理
  fmt.Print(err)
}

if err := doError(); err != nil {
  // エラー時の処理
  fmt.Print(err)
}

func doError() error{
  return errors.New("error")
}
```

### メソッドの可変長引数
```
func sum(n ...int) int {
	fmt.Println(n)

	sum := 0
	for i, v := range n {
		sum += v
		fmt.Println("インデックス:", i, "値:", v, "・・・現在の合計:", sum)
	}
	return sum
}

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("引数で渡した数の総和:", x)
}

```

### 複数の引数

```
func f1() (int, string) {
	return 999, "hoge"
}

func main() {
	var a, b = f1()
	fmt.Println(a)
	fmt.Println(b)
}
```

## go run
コンパイル型なのでソース段階で動かすためには
```
go run main.go

```
## go build
バイナリファイルの作成
```
go build -o ファイル名　

#通常のシェルスクリプトの実行のように動かす
./ファイル名
```


## プラグイン
- go-plus(atom)
  保存時の自動コンパイル、テストはうざいので設定で切ってもいいかも

## 環境変数関連
```go env``` で見れる。特にGO_ROOT,GO_PATHに注意。

## 並列直列に関して
参考URL https://qiita.com/suin/items/82ecb6f63ff4104d4f5d
- parallel
    - chokuretsu 直列処理(比較のためで新ネタはなし)
    - heiretsu.go 並列処理　プレーンなもので無名関数使用
    - heiretsu2.go 並列処理　heiretsu1と同じだがチェネルの終了にfor文
    - heiretsu3.go 並列処理　heiretsu1と同じだが書き方を修正

    参考URL https://qiita.com/TakaakiFuruse/items/241578174fd2f00aaa8a
    - heiretsu4.go 並列処理　書き方を若干変更(無名関数を使わないためこちらのほうが一般的かも・・・)
    - heiretsu5.go 並列処理　同一処理を並行。これが一般的な使い方かと・・

    他参考URL https://tikasan.hatenablog.com/entry/2017/05/20/170050

## テスト、任意ファイルの実行について

通常はmain.goから実行し、任意のファイルを実行することができないので必須

<br>
別のライブラリを入れたりなどは必要なし

- sample
    - hello.go
    - hello_test.go
    同一階層かつ_testをつけることでtestスクリプトになる

テストコマンド
```
#対象のディレクトリに移動している必要がある
cd ./sample
#全テスト実行
go test  -run ''
#特定文字列を含むメソッドのみ
go test  -run Dummy
```
## interface型

sample/hello.goのsampleInterface


## インポート(import)に関して

- 標準のライブラリ
- 相対パス
- GO_PATHのものなど(主にライブラリ)

## パッケージ
http://cuto.unirita.co.jp/gostudy/post/go-package/

ルール
- 同一階層ではpackageを統一する必要がある(例外はテストコード)
- フォルダ名とあっていなくてもいい    

## 開発環境(エディタ)メモ
- dockerだけだと、エディタで構文チェックなどが使えない。エディタで色々な機能(主に構文チェックや未使用モジュールの強制処理)を走らせるために、ホスト機にgoのインストールが必要

一括で入れておくと良いライブラリなど($HOME/go配下にライブラリあり)
- gocode
- gopkgs
- go-outline
- gocode-gomod
- godef
- goreturns
- golint 静的解析ツール

## 環境変数の読み込み
source load_environment
